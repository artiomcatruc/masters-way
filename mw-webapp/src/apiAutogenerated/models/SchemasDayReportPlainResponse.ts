// @ts-nocheck
/* eslint-disable */
/**
 * Masters way API
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * The version of the OpenAPI document: 1.0
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import type { SchemasCommentPlainResponse } from './SchemasCommentPlainResponse';
import {
    SchemasCommentPlainResponseFromJSON,
    SchemasCommentPlainResponseFromJSONTyped,
    SchemasCommentPlainResponseToJSON,
} from './SchemasCommentPlainResponse';
import type { SchemasJobDonePlainResponse } from './SchemasJobDonePlainResponse';
import {
    SchemasJobDonePlainResponseFromJSON,
    SchemasJobDonePlainResponseFromJSONTyped,
    SchemasJobDonePlainResponseToJSON,
} from './SchemasJobDonePlainResponse';
import type { SchemasPlanPlainResponse } from './SchemasPlanPlainResponse';
import {
    SchemasPlanPlainResponseFromJSON,
    SchemasPlanPlainResponseFromJSONTyped,
    SchemasPlanPlainResponseToJSON,
} from './SchemasPlanPlainResponse';
import type { SchemasProblemPlainResponse } from './SchemasProblemPlainResponse';
import {
    SchemasProblemPlainResponseFromJSON,
    SchemasProblemPlainResponseFromJSONTyped,
    SchemasProblemPlainResponseToJSON,
} from './SchemasProblemPlainResponse';

/**
 * 
 * @export
 * @interface SchemasDayReportPlainResponse
 */
export interface SchemasDayReportPlainResponse {
    /**
     * 
     * @type {Array<SchemasCommentPlainResponse>}
     * @memberof SchemasDayReportPlainResponse
     */
    comments?: Array<SchemasCommentPlainResponse>;
    /**
     * 
     * @type {string}
     * @memberof SchemasDayReportPlainResponse
     */
    createdAt?: string;
    /**
     * 
     * @type {boolean}
     * @memberof SchemasDayReportPlainResponse
     */
    isDayOff?: boolean;
    /**
     * 
     * @type {Array<SchemasJobDonePlainResponse>}
     * @memberof SchemasDayReportPlainResponse
     */
    jobsDone?: Array<SchemasJobDonePlainResponse>;
    /**
     * 
     * @type {Array<SchemasPlanPlainResponse>}
     * @memberof SchemasDayReportPlainResponse
     */
    plans?: Array<SchemasPlanPlainResponse>;
    /**
     * 
     * @type {Array<SchemasProblemPlainResponse>}
     * @memberof SchemasDayReportPlainResponse
     */
    problems?: Array<SchemasProblemPlainResponse>;
    /**
     * 
     * @type {string}
     * @memberof SchemasDayReportPlainResponse
     */
    updatedAt?: string;
}

/**
 * Check if a given object implements the SchemasDayReportPlainResponse interface.
 */
export function instanceOfSchemasDayReportPlainResponse(
    value: object
): boolean {
    let isInstance = true;

    return isInstance;
}

export function SchemasDayReportPlainResponseFromJSON(json: any): SchemasDayReportPlainResponse {
    return SchemasDayReportPlainResponseFromJSONTyped(json, false);
}

export function SchemasDayReportPlainResponseFromJSONTyped(
    json: any,
    ignoreDiscriminator: boolean
): SchemasDayReportPlainResponse {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'comments': !exists(json, 'comments') ? undefined : ((json['comments'] as Array<any>).map(SchemasCommentPlainResponseFromJSON)),
        'createdAt': !exists(json, 'createdAt') ? undefined : json['createdAt'],
        'isDayOff': !exists(json, 'isDayOff') ? undefined : json['isDayOff'],
        'jobsDone': !exists(json, 'jobsDone') ? undefined : ((json['jobsDone'] as Array<any>).map(SchemasJobDonePlainResponseFromJSON)),
        'plans': !exists(json, 'plans') ? undefined : ((json['plans'] as Array<any>).map(SchemasPlanPlainResponseFromJSON)),
        'problems': !exists(json, 'problems') ? undefined : ((json['problems'] as Array<any>).map(SchemasProblemPlainResponseFromJSON)),
        'updatedAt': !exists(json, 'updatedAt') ? undefined : json['updatedAt'],
    };
}


export function SchemasDayReportPlainResponseToJSON(value?: SchemasDayReportPlainResponse | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'comments': value.comments === undefined ? undefined : ((value.comments as Array<any>).map(SchemasCommentPlainResponseToJSON)),
        'createdAt': value.createdAt,
        'isDayOff': value.isDayOff,
        'jobsDone': value.jobsDone === undefined ? undefined : ((value.jobsDone as Array<any>).map(SchemasJobDonePlainResponseToJSON)),
        'plans': value.plans === undefined ? undefined : ((value.plans as Array<any>).map(SchemasPlanPlainResponseToJSON)),
        'problems': value.problems === undefined ? undefined : ((value.problems as Array<any>).map(SchemasProblemPlainResponseToJSON)),
        'updatedAt': value.updatedAt,
    };
}
