import { OnInit } from "@angular/core";

enum LoginModes {
  Kubeconfig = 'kubeconfig',
  Basic = 'basic',
  Token = 'token',
}

export class LoginComponent implements OnInit {
  loginModes = LoginModes;

  errors: KdError[] = [];

  private enablesAuthenticationModes_: AuthenticationMode[] = [];

  constructor(
    private readonly http_: HttpClient,
  ) {}

  ngOnInit(): void {
    this.http_
      .get<EnabledAuthenticationModes>('api/v1/login/modes')
      .subscribe((enabledMode: EnableAuthenticationModes) => {
        this.enablesAuthenticationModes_ = enabledMode.modes;
        this.enablesAuthenticationModes_.push(LoginModes.Kubeconfig);
        this.selectedAuthenticationMode as LoginModels)
      })
  }
}
