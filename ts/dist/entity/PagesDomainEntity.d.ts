import { CloudflarePagesEntityBase } from '../CloudflarePagesEntityBase';
import type { CloudflarePagesSDK } from '../CloudflarePagesSDK';
import type { Control } from '../types';
import type { PagesDomain, PagesDomainLoadMatch, PagesDomainListMatch, PagesDomainCreateData, PagesDomainUpdateData, PagesDomainRemoveMatch } from '../CloudflarePagesTypes';
declare class PagesDomainEntity extends CloudflarePagesEntityBase<PagesDomain> {
    constructor(client: CloudflarePagesSDK, entopts: any);
    make(this: PagesDomainEntity): PagesDomainEntity;
    load(this: any, reqmatch?: PagesDomainLoadMatch, ctrl?: Control): Promise<PagesDomainEntity>;
    list(this: any, reqmatch?: PagesDomainListMatch, ctrl?: Control): Promise<PagesDomainEntity[]>;
    create(this: any, reqdata?: PagesDomainCreateData, ctrl?: Control): Promise<PagesDomainEntity>;
    update(this: any, reqdata?: PagesDomainUpdateData, ctrl?: Control): Promise<PagesDomainEntity>;
    remove(this: any, reqmatch?: PagesDomainRemoveMatch, ctrl?: Control): Promise<PagesDomainEntity>;
}
export { PagesDomainEntity };
