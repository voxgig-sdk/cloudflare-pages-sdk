import { CloudflarePagesEntityBase } from '../CloudflarePagesEntityBase';
import type { CloudflarePagesSDK } from '../CloudflarePagesSDK';
import type { Control } from '../types';
import type { PagesProject, PagesProjectLoadMatch, PagesProjectListMatch, PagesProjectCreateData, PagesProjectUpdateData, PagesProjectRemoveMatch } from '../CloudflarePagesTypes';
declare class PagesProjectEntity extends CloudflarePagesEntityBase<PagesProject> {
    constructor(client: CloudflarePagesSDK, entopts: any);
    make(this: PagesProjectEntity): PagesProjectEntity;
    load(this: any, reqmatch?: PagesProjectLoadMatch, ctrl?: Control): Promise<PagesProjectEntity>;
    list(this: any, reqmatch?: PagesProjectListMatch, ctrl?: Control): Promise<PagesProjectEntity[]>;
    create(this: any, reqdata?: PagesProjectCreateData, ctrl?: Control): Promise<PagesProjectEntity>;
    update(this: any, reqdata?: PagesProjectUpdateData, ctrl?: Control): Promise<PagesProjectEntity>;
    remove(this: any, reqmatch?: PagesProjectRemoveMatch, ctrl?: Control): Promise<PagesProjectEntity>;
}
export { PagesProjectEntity };
