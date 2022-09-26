import {Component, Inject, Input} from '@angular/core';
import {IMessage} from '@api/root.ui';
import {MESSAGES_DI_TOKEN} from '../../../index.messages';
import {Animations} from '../../animations/animations';

@Component({
  selector: 'kd-card',
  templateUrl: './template.html',
  styleUrls: ['./style.scss'],
  animations: [Animations.expandInOut],
})
export class CardComponent {
  @Input() initialized = true;
  @Input() role: 'inner' | 'table' | 'inner-content';
  @Input() expandable = true;
  @Input() expanded = true;
  private classes_: string[] = [];

  @Input()
  set titleClasses(val: string) {
    this.classes_ = val.split(/\s+/);
  }

  constructor(@Inject(MESSAGES_DI_TOKEN) readonly messgae: IMessage) {}

  expand(): void {
    if (this.expandable) {
      this.expanded = !this.expanded;
    }
  }

  onCardHeaderClick(): void {
    if (this.expandable && !this.expanded) {
      this.expanded = true;
    }
  }

  getTitleClasses(): {[clsName: string]: boolean} {
    const ngCls = {} as {[clsName: string]: boolean};
    if (!this.expanded) {
      ngCls['kd-minimized-card-header'] = true;
    }

    if (this.expandable) {
      ngCls['kd-card-header'] = true;
    }

    for (const cls of this.classes_) {
      ngCls[cls] = true;
    }

    return ngCls;
  }
}