# Typed models for the CloudflarePages SDK.
#
# GENERATED from the API model: main.kit.entity.<e>.fields{} and per-op
# params (op.<name>.points[].g.params[]). Field/param types come from the
# canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
# @voxgig/apidef VALID_CANON). Do not edit by hand.
#
# These are TypedDicts, not dataclasses: the SDK ops return/accept plain dicts
# at runtime, and a TypedDict IS a dict shape, so the types match the runtime.
# Optional (req:false) keys are modelled as TypedDict key-optionality
# (total=False), split into a required base + total=False subclass when a type
# has both required and optional keys.

from __future__ import annotations

from typing import TypedDict, Any


class PagesDeploymentRequired(TypedDict):
    errors: list
    messages: list
    result: dict
    success: bool


class PagesDeployment(PagesDeploymentRequired, total=False):
    filters: list
    id: str


class PagesDeploymentLoadMatch(TypedDict):
    account_id: str
    id: str
    project_id: str


class PagesDeploymentListMatchRequired(TypedDict):
    account_id: str


class PagesDeploymentListMatch(PagesDeploymentListMatchRequired, total=False):
    project_name: str
    env: str
    page: int
    per_page: int
    deployment_id: str
    project_id: str


class PagesDeploymentCreateDataRequired(TypedDict):
    account_id: str
    errors: list
    messages: list
    result: dict
    success: bool


class PagesDeploymentCreateData(PagesDeploymentCreateDataRequired, total=False):
    deployment_id: str
    project_id: str
    project_name: str
    filters: list
    id: str


class PagesDeploymentRemoveMatchRequired(TypedDict):
    account_id: str
    project_id: str


class PagesDeploymentRemoveMatch(PagesDeploymentRemoveMatchRequired, total=False):
    deployment_id: str
    tail_id: str
    id: str
    force: bool


class PagesDomainRequired(TypedDict):
    errors: list
    messages: list
    name: str
    result: dict
    success: bool


class PagesDomain(PagesDomainRequired, total=False):
    id: str


class PagesDomainLoadMatch(TypedDict):
    account_id: str
    id: str
    project_id: str


class PagesDomainListMatch(TypedDict):
    account_id: str
    project_name: str


class PagesDomainCreateDataRequired(TypedDict):
    account_id: str
    project_name: str
    errors: list
    messages: list
    name: str
    result: dict
    success: bool


class PagesDomainCreateData(PagesDomainCreateDataRequired, total=False):
    id: str


class PagesDomainUpdateDataRequired(TypedDict):
    account_id: str
    id: str
    project_id: str


class PagesDomainUpdateData(PagesDomainUpdateDataRequired, total=False):
    errors: list
    messages: list
    name: str
    result: dict
    success: bool


class PagesDomainRemoveMatch(TypedDict):
    account_id: str
    id: str
    project_id: str


class PagesPagesAssetsCheckMissing(TypedDict):
    hashes: list


class PagesPagesAssetsCheckMissingCreateData(TypedDict):
    hashes: list


class PagesProjectRequired(TypedDict):
    errors: list
    messages: list
    name: str
    production_branch: str
    result: dict
    source: dict
    success: bool


class PagesProject(PagesProjectRequired, total=False):
    build_config: dict
    deployment_configs: dict
    id: str


class PagesProjectLoadMatch(TypedDict):
    account_id: str
    id: str


class PagesProjectListMatchRequired(TypedDict):
    account_id: str


class PagesProjectListMatch(PagesProjectListMatchRequired, total=False):
    page: int
    per_page: int
    project_name: str


class PagesProjectCreateDataRequired(TypedDict):
    account_id: str
    errors: list
    messages: list
    name: str
    production_branch: str
    result: dict
    source: dict
    success: bool


class PagesProjectCreateData(PagesProjectCreateDataRequired, total=False):
    build_config: dict
    deployment_configs: dict
    id: str


class PagesProjectUpdateDataRequired(TypedDict):
    account_id: str
    id: str


class PagesProjectUpdateData(PagesProjectUpdateDataRequired, total=False):
    build_config: dict
    deployment_configs: dict
    errors: list
    messages: list
    name: str
    production_branch: str
    result: dict
    source: dict
    success: bool


class PagesProjectRemoveMatch(TypedDict):
    account_id: str
    id: str


class Upload(TypedDict):
    errors: list
    messages: list
    success: bool


class UploadCreateData(TypedDict):
    errors: list
    messages: list
    success: bool


class UpsertHash(TypedDict):
    errors: list
    hashes: list
    messages: list
    success: bool


class UpsertHashCreateData(TypedDict):
    errors: list
    hashes: list
    messages: list
    success: bool
