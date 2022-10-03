import {HttpClient, HttpErrorResponse} from '@angular/common/http';
import { OnInit } from "@angular/core";

import {AuthenticationMode} from '@api/root.api';

enum LoginModes {
  Kubeconfig = 'kubeconfig',
  Basic = 'basic',
  Token = 'token',
}

export class LoginComponent implements OnInit {
  loginModes = LoginModes;

  errors: KdError[] = [];

  private enabledAuthenticationModes_: AuthenticationMode[] = [];
  private isLoginSkippable_ = false;
  private kubeconfig_: string;
  private token_: string;
  private username_: string;
  private password_: string;

  constructor(
    private readonly authService_: Auservce
    private readonly http_: HttpClient,
  ) {}

  ngOnInit(): void {
    this.http_
      .get<EnabledAuthenticationModes>('api/v1/login/modes')
      .subscribe((enabledMode: EnableAuthenticationModes) => {
        this.enablesAuthenticationModes_ = enabledMode.modes;
        this.enablesAuthenticationModes_.push(LoginModes.Kubeconfig);
        this.selectedAuthenticationMode = this.selectedAuthenticationMode
          ? (this.selectedAuthenticationMode as LoginModes)
          : (this.enablesAuthenticationModes_[0] as LoginModes)
      });

    this.http_
      .set
  }
}
