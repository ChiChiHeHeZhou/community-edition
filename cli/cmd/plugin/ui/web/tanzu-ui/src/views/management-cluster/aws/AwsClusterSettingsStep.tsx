// React imports
import React, { ChangeEvent, useContext, useState, useEffect } from 'react';

// Library imports
import { ClarityIcons, blockIcon, blocksGroupIcon, clusterIcon } from '@cds/core/icon';
import { CdsControlMessage } from '@cds/react/forms';
import { FieldError, useForm } from 'react-hook-form';
import { CdsIcon } from '@cds/react/icon';
import { CdsButton } from '@cds/react/button';
import { CdsSelect } from '@cds/react/select';

// App imports
import { INPUT_CHANGE } from '../../../state-management/actions/Form.actions';
import { StepProps } from '../../../shared/components/wizard/Wizard';
import { STORE_SECTION_FORM } from '../../../state-management/reducers/Form.reducer';
import {
    NodeProfile,
    NodeInstanceType,
    nodeInstanceTypeValidation,
} from '../../../shared/components/FormInputComponents/NodeProfile/NodeProfile';
import { ClusterName, clusterNameValidation } from '../../../shared/components/FormInputComponents/ClusterName/ClusterName';
import { AwsStore } from '../../../state-management/stores/Store.aws';
import { yupResolver } from '@hookform/resolvers/yup/dist/yup';
import * as yup from 'yup';
import { AwsService, AWSVirtualMachine } from '../../../swagger-api';
import RetrieveOSImages from '../../../shared/components/FormInputComponents/RetrieveOSImages/RetrieveOSImages';

ClarityIcons.addIcons(blockIcon, blocksGroupIcon, clusterIcon);

type AWS_CLUSTER_SETTING_STEP_FIELDS = 'NODE_PROFILE' | 'IMAGE_INFO' | 'CLUSTER_NAME';

interface MCSettings extends StepProps {
    message?: string;
    deploy: () => void;
    defaultData?: { [key: string]: any };
}

const nodeInstanceTypes: NodeInstanceType[] = [
    {
        id: 'SINGLE_NODE',
        label: 'Single node',
        icon: 'block',
        description: 'Create a single control plane node with a medium instance type',
    },
    {
        id: 'HIGH_AVAILABILITY',
        label: 'High availability',
        icon: 'blocks-group',
        description: 'Create a multi-node control plane with a medium instance type',
    },
    {
        id: 'PRODUCTION_READY',
        label: 'Production-ready (High availability)',
        icon: 'blocks-group',
        isSolidIcon: true,
        description: 'Create a multi-node control plane with a large instance type',
    },
];

interface AwsClusterSettingFormInputs {
    CLUSTER_NAME: string;
    NODE_PROFILE: string;
    IMAGE_INFO: string;
}

function createYupSchemaObject() {
    return {
        IMAGE_INFO: yupStringRequired('Please select an OS image'),
        NODE_PROFILE: nodeInstanceTypeValidation(),
        CLUSTER_NAME: clusterNameValidation(),
    };
}

function yupStringRequired(errorMessage: string) {
    return yup.string().nullable().required(errorMessage);
}

function AwsClusterSettingsStep(props: Partial<MCSettings>) {
    const { handleValueChange, currentStep, deploy, defaultData, message } = props;
    const { awsState } = useContext(AwsStore);
    const awsClusterSettingsFormSchema = yup.object(createYupSchemaObject()).required();
    const methods = useForm<AwsClusterSettingFormInputs>({
        resolver: yupResolver(awsClusterSettingsFormSchema),
    });
    const [images, setImages] = useState<AWSVirtualMachine[]>([]);
    const {
        handleSubmit,
        formState: { errors },
        register,
        setValue,
    } = methods;

    const setImageParameters = (image) => {
        if (handleValueChange) {
            handleValueChange(INPUT_CHANGE, 'IMAGE_INFO', image, currentStep, errors);
        }
    };

    const handleMCCreation = () => {
        if (deploy) {
            deploy();
        }
    };

    useEffect(() => {
        AwsService.getAwsosImages(awsState[STORE_SECTION_FORM].REGION).then((data) => {
            setImages(data);
            setImageParameters(data[0]);
        });
    }, [awsState[STORE_SECTION_FORM].REGION]);

    let initialSelectedInstanceTypeId = awsState[STORE_SECTION_FORM].NODE_PROFILE;

    if (!initialSelectedInstanceTypeId) {
        initialSelectedInstanceTypeId = nodeInstanceTypes[0].id;
        setValue('NODE_PROFILE', initialSelectedInstanceTypeId);
    }
    const [selectedInstanceTypeId, setSelectedInstanceTypeId] = useState(initialSelectedInstanceTypeId);

    const onFieldChange = (data: string, field: AWS_CLUSTER_SETTING_STEP_FIELDS) => {
        if (handleValueChange) {
            handleValueChange(INPUT_CHANGE, field, data, currentStep, errors);
            setValue(field, data, { shouldValidate: true });
        }
    };

    const onInstanceTypeChange = (instanceType: string) => {
        onFieldChange(instanceType, 'NODE_PROFILE');
        setSelectedInstanceTypeId(instanceType);
    };

    const onClusterNameChange = (clusterName: string) => {
        onFieldChange(clusterName, 'CLUSTER_NAME');
    };

    const onOsImageSelected = (clusterName: string) => {
        onFieldChange(clusterName, 'IMAGE_INFO');
    };

    return (
        <div className="cluster-settings-container" cds-layout="m:lg">
            <h3>Management Cluster settings</h3>
            <div cds-layout="grid gap:m" key="section-holder">
                <div cds-layout="col:6" key="cluster-name-section">
                    <ClusterName
                        field={'CLUSTER_NAME'}
                        errors={errors}
                        register={register}
                        clusterNameChange={onClusterNameChange}
                        placeholderClusterName={'my-aws-cluster'}
                    />
                </div>
                <div cds-layout="col:6" key="instance-type-section">
                    <NodeProfile
                        field={'NODE_PROFILE'}
                        nodeInstanceTypes={nodeInstanceTypes}
                        errors={errors}
                        register={register}
                        nodeInstanceTypeChange={onInstanceTypeChange}
                        selectedInstanceId={selectedInstanceTypeId}
                    />
                </div>

                {/* <div cds-layout="col:12">{AwsOSImageSection('IMAGE_INFO', images, errors, register, onFieldChange)}</div> */}

                <div cds-layout="col:12">
                    <RetrieveOSImages
                        osImageTitle={'Amazon Machine Image(AMI)'}
                        images={images}
                        field={'IMAGE_INFO'}
                        errors={errors}
                        register={register}
                        onOsImageSelected={onOsImageSelected}
                    />
                </div>

                <div cds-layout="grid col:12 p-t:lg">
                    <CdsButton cds-layout="col:start-1" status="success" onClick={handleSubmit(handleMCCreation)}>
                        <CdsIcon shape="cluster" size="sm"></CdsIcon>
                        Create Management cluster
                    </CdsButton>
                    <CdsButton cds-layout="col:end-12" action="flat">
                        View configuration details
                    </CdsButton>
                </div>

                {/* <div cds-layout="col:8">
                    <h1>{'Amazon Machine Image(AMI)'}</h1>
                    <CdsSelect layout="compact">
                        <label>OS Image with Kubernetes </label>
                        <select className="select-sm-width" data-testid="image-select" onChange={handleMachineImageChange}>
                            {images.map((image) => (
                                <option key={image.name} value={image.name}>
                                    {image.name}
                                </option>
                            ))}
                        </select>
                    </CdsSelect>
                </div> */}
            </div>
        </div>
    );
}

function AwsOSImageSection(
    field: AWS_CLUSTER_SETTING_STEP_FIELDS,
    vmTemplates: AWSVirtualMachine[],
    errors: { [key: string]: FieldError | undefined },
    register: any,
    onOsImageSelected: (osImage: string, field: AWS_CLUSTER_SETTING_STEP_FIELDS) => void
) {
    const handleOsImageSelect = (event: ChangeEvent<HTMLSelectElement>) => {
        onOsImageSelected(event.target.value || '', field);
    };
    const fieldError = errors[field];
    return (
        <div cds-layout="m:lg">
            <h1>{'Amazon Machine Image(AMI)'}</h1>
            <CdsSelect layout="compact" controlWidth="shrink">
                <label>OS Image with Kubernetes </label>
                <select {...register('IMAGE_INFO')} onChange={handleOsImageSelect}>
                    {AwsOSDropdownOptions(vmTemplates)}
                </select>
            </CdsSelect>
            {fieldError && (
                <div>
                    &nbsp;<CdsControlMessage status="error">{fieldError.message}</CdsControlMessage>
                </div>
            )}
        </div>
    );
}

function AwsOSDropdownOptions(images: AWSVirtualMachine[]) {
    // if (vmTemplates && vmTemplates.length === 1) {
    //     return (

    //         <option key={vmTemplates[0].name} value={vmTemplates[0].name}>
    //             {vmTemplates[0].name}
    //         </option>
    //     );
    // }
    return (
        <>
            {images.map((dc) => (
                <option key={dc.name}>{dc.name}</option>
            ))}
        </>
    );
}

export default AwsClusterSettingsStep;
