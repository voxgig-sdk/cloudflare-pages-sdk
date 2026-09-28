import { CloudflarePagesEntityBase } from '../CloudflarePagesEntityBase';
import type { CloudflarePagesSDK } from '../CloudflarePagesSDK';
import type { Control } from '../types';
import type { UpsertHash, UpsertHashCreateData } from '../CloudflarePagesTypes';
declare class UpsertHashEntity extends CloudflarePagesEntityBase<UpsertHash> {
    constructor(client: CloudflarePagesSDK, entopts: any);
    make(this: UpsertHashEntity): UpsertHashEntity;
    create(this: any, reqdata?: UpsertHashCreateData, ctrl?: Control): Promise<UpsertHashEntity>;
}
export { UpsertHashEntity };
