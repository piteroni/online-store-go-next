/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { CancelablePromise } from "../core/CancelablePromise"
import type { BaseHttpRequest } from "../core/BaseHttpRequest"

export class DefaultService {
  constructor(public readonly httpRequest: BaseHttpRequest) {}

  /**
   * Login With Admin
   * post
   * @param body
   * @returns any
   * @throws ApiError
   */
  public postAdminLogin(body: {
    loginId: string
    password: string
  }): CancelablePromise<{
    token: string
  }> {
    return this.httpRequest.request({
      method: "POST",
      url: "/admin/login",
      body: body,
      errors: {
        400: `Bad Request`,
      },
    })
  }
}
