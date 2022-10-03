import {HttpClient, HttpHeaders} from '@angular/common/http';
import { Inject } from '@angular/core';
import { IConfig } from '@api/root.ui';
import {CookieService} from 'ngx-cookie-service';
import { Observable } from 'rxjs';
import {switchMap, take} from 'rxjs/operators';
import {CsrfToken, LoginStatus} from 'typings/root.api';
import {CONFIG_DI_TOKEN} from '../../../index.config';

import {CsrfTokenService} from './csrftoken';
import {KdStateService} from './state';

export class AuthService {
  constructor(
    private readonly cookies_: CookieService,
    private readonly http_: HttpClient,
    private readonly csrfTokenService_: CsrfTokenService,
    private readonly stateService_: KdStateService,
    @Inject(CONFIG_DI_TOKEN) private readonly config_: IConfig
  ) {
    this.init_();
  }

  private init_() {
    this.stateService_.onBefore.pipe(switchMap(() => this.getLoginStatus())).subscribe(status => {
      if (this.isAuthenticationEnabled(status)) {
        this
      }
    });
  }

  private getTokenCookie_(): string {
    return this.cookies_.get(this.config_.authTokenCookieName) || '';
  }

  refreshToken(): void {
    const token = this.getTokenCookie_();
    if (token.length === 0) return;

    this.csrfTokenService_
      .getTokenForAction('token')
      .pipe(
        switchMap(csrfToken => {
          return this.http_.post<
        })
      )
  }

  /**
   * Checks authentication is enabled. It is enabled only on HTTPS. Can be
   * overridden by 'enable-insecure-login' flag passed to dashboard.
   */
  isAuthenticationEnabled(loginStatus: LoginStatus): boolean {
    return loginStatus.httpsMode;
  }

  getLoginStatus(): Observable<LoginStatus> {
    return this.http_.get<LoginStatus>('api/v1/login/status');
  }
}