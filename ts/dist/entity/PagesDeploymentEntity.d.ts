import { CloudflarePagesEntityBase } from '../CloudflarePagesEntityBase';
import type { CloudflarePagesSDK } from '../CloudflarePagesSDK';
import type { Control } from '../types';
import type { PagesDeployment, PagesDeploymentLoadMatch, PagesDeploymentListMatch, PagesDeploymentCreateData, PagesDeploymentRemoveMatch } from '../CloudflarePagesTypes';
declare class PagesDeploymentEntity extends CloudflarePagesEntityBase<PagesDeployment> {
    constructor(client: CloudflarePagesSDK, entopts: any);
    make(this: PagesDeploymentEntity): PagesDeploymentEntity;
    load(this: any, reqmatch?: PagesDeploymentLoadMatch, ctrl?: Control): Promise<PagesDeploymentEntity>;
    list(this: any, reqmatch?: PagesDeploymentListMatch, ctrl?: Control): Promise<PagesDeploymentEntity[]>;
    create(this: any, reqdata?: PagesDeploymentCreateData, ctrl?: Control): Promise<PagesDeploymentEntity>;
    remove(this: any, reqmatch?: PagesDeploymentRemoveMatch, ctrl?: Control): Promise<PagesDeploymentEntity>;
}
export { PagesDeploymentEntity };
