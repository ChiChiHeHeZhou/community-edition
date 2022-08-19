// App imports
import { AWSVirtualMachine } from '../../../../../swagger-api';
import { AWSKeyPair } from '../../../../../swagger-api/models/AWSKeyPair';
import { validateDefaultNodeType } from '../../../../../shared/constants/defaults/aws.defaults';
export class AwsDefaults {
    // The strategy of deciding default os image
    static selectDefalutOsImage = (osImages: AWSVirtualMachine[]) => {
        return osImages && osImages.length > 0 ? osImages[0] : undefined;
    };

    static selectDefalutEC2KeyPairs = (keyPairs: AWSKeyPair[]) => {
        return keyPairs && keyPairs.length > 0 ? keyPairs[0].name : undefined;
    };

    static setDefaultNodeType = (nodeTypeList: string[], nodeProfile: string) => {
        if (nodeTypeList.indexOf(validateDefaultNodeType(nodeProfile)) > -1) {
            return validateDefaultNodeType(nodeProfile);
        } else {
            // TODO: refactor to select optimal nodeType when preferred default not found
            return nodeTypeList[nodeTypeList.length / 2];
        }
    };
}
