import { Component, Inject } from '@angular/core';
import {VersionInfo} from '@api/root.ui';
import {AssetsService} from '@common/services/global/assets';
import {ConfigService} from '@common/services/global/config';

@Component({
  selector: 'kd-about',
  templateUrl: './template.html',
  styleUrls: ['./style.scss'],
})
export class AboutComponent {
  lastestCopyrightYear: number;
  versionInfo: VersionInfo;

  constructor(@Inject(AssetsService) public assets: AssetsService, config: ConfigService) {
    this.versionInfo = config.getVersionInfo();
    this.lastestCopyrightYear = new Date().getFullYear();
  }
}