import { PagesDeploymentEntity } from './entity/PagesDeploymentEntity';
import { PagesDomainEntity } from './entity/PagesDomainEntity';
import { PagesPagesAssetsCheckMissingEntity } from './entity/PagesPagesAssetsCheckMissingEntity';
import { PagesProjectEntity } from './entity/PagesProjectEntity';
import { UploadEntity } from './entity/UploadEntity';
import { UpsertHashEntity } from './entity/UpsertHashEntity';
export type * from './CloudflarePagesTypes';
import { inspect } from 'node:util';
import type { Context, Feature } from './types';
import { config } from './Config';
import { CloudflarePagesEntityBase } from './CloudflarePagesEntityBase';
import { Utility } from './utility/Utility';
import { BaseFeature } from './feature/base/BaseFeature';
declare const stdutil: Utility;
declare class CloudflarePagesSDK {
    _mode: string;
    _options: any;
    _utility: Utility;
    _features: Feature[];
    _rootctx: Context;
    constructor(options?: any);
    options(): any;
    utility(): any;
    prepare(fetchargs?: any): Promise<any>;
    direct(fetchargs?: any): Promise<Error | {
        ok: boolean;
        status: number;
        headers: any;
        data: any;
        err?: undefined;
    } | {
        ok: boolean;
        err: any;
        status?: undefined;
        headers?: undefined;
        data?: undefined;
    }>;
    _rawRequest(fetchargs?: any): Promise<Error | {
        ok: boolean;
        status: number;
        headers: any;
        data: any;
        err?: undefined;
    } | {
        ok: boolean;
        err: any;
        status?: undefined;
        headers?: undefined;
        data?: undefined;
    }>;
    graphql(query: string, variables?: any, ctrl?: any): Promise<any>;
    PagesDeployment(entopts?: Record<string, any>): PagesDeploymentEntity;
    PagesDomain(entopts?: Record<string, any>): PagesDomainEntity;
    PagesPagesAssetsCheckMissing(entopts?: Record<string, any>): PagesPagesAssetsCheckMissingEntity;
    PagesProject(entopts?: Record<string, any>): PagesProjectEntity;
    Upload(entopts?: Record<string, any>): UploadEntity;
    UpsertHash(entopts?: Record<string, any>): UpsertHashEntity;
    static test(testoptsarg?: any, sdkoptsarg?: any): CloudflarePagesSDK;
    tester(testopts?: any, sdkopts?: any): CloudflarePagesSDK;
    toJSON(): {
        name: string;
    };
    toString(): string;
    [inspect.custom](): string;
}
declare const SDK: typeof CloudflarePagesSDK;
export { stdutil, config, BaseFeature, CloudflarePagesEntityBase, CloudflarePagesSDK, SDK, };
