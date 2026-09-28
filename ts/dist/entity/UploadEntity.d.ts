import { CloudflarePagesEntityBase } from '../CloudflarePagesEntityBase';
import type { CloudflarePagesSDK } from '../CloudflarePagesSDK';
import type { Control } from '../types';
import type { Upload, UploadCreateData } from '../CloudflarePagesTypes';
declare class UploadEntity extends CloudflarePagesEntityBase<Upload> {
    constructor(client: CloudflarePagesSDK, entopts: any);
    make(this: UploadEntity): UploadEntity;
    create(this: any, reqdata?: UploadCreateData, ctrl?: Control): Promise<UploadEntity>;
}
export { UploadEntity };
