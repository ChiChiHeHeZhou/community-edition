// App imports
import AwsLogo from '../../assets/aws.svg';
import AzureLogo from '../../assets/azure.svg';
import DockerLogo from '../../assets/docker.svg';
import VSphereLogo from '../../assets/vsphere.svg';

export interface ProviderInfo {
    [key: string]: ProviderData;
}

export interface ProviderData {
    name: string;
    logo: string;
}

const providerInfo: ProviderInfo = {
    aws: {
        name: 'AWS',
        logo: AwsLogo,
    },
    azure: {
        name: 'Azure',
        logo: AzureLogo,
    },
    docker: {
        name: 'Docker',
        logo: DockerLogo,
    },
    vsphere: {
        name: 'vSphere',
        logo: VSphereLogo,
    },
};

/**
 * @method retrieveProviderInfo
 * @param providerName - provider name string to reference in providerInfo map
 * Returns human-readable provider name and logo
 */
export function retrieveProviderInfo(providerName: string): ProviderData {
    if (!providerName || !providerInfo[providerName]) {
        console.log(`retrieveProviderInfo() was called with an invalid provider name: "${providerName}"`);
    }

    return providerInfo[providerName];
}