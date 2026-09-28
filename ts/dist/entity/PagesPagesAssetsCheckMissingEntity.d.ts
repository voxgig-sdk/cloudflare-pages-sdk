import { CloudflarePagesEntityBase } from '../CloudflarePagesEntityBase';
import type { CloudflarePagesSDK } from '../CloudflarePagesSDK';
import type { Control } from '../types';
import type { PagesPagesAssetsCheckMissing, PagesPagesAssetsCheckMissingCreateData } from '../CloudflarePagesTypes';
declare class PagesPagesAssetsCheckMissingEntity extends CloudflarePagesEntityBase<PagesPagesAssetsCheckMissing> {
    constructor(client: CloudflarePagesSDK, entopts: any);
    make(this: PagesPagesAssetsCheckMissingEntity): PagesPagesAssetsCheckMissingEntity;
    create(this: any, reqdata?: PagesPagesAssetsCheckMissingCreateData, ctrl?: Control): Promise<PagesPagesAssetsCheckMissingEntity>;
}
export { PagesPagesAssetsCheckMissingEntity };
