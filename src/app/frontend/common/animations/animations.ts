import { animate, AUTO_STYLE, state, style, transition, trigger } from "@angular/animations";

const DEFAULT_TRANSITION_TIME = '500ms ease-in-out';

export class Animations {
  static easeOut = trigger('easeOut', [
    transition(`${AUTO_STYLE} => void`, [style({opacity: 1}), animate(DEFAULT_TRANSITION_TIME, style({opacity: 0}))]),
  ]);

  static easeInOut = trigger('easeInOut', [
    transition(`void => ${AUTO_STYLE}`, [style({opacity: 0}), animate(DEFAULT_TRANSITION_TIME, style({opacity: 1}))]),
  ]);

  static expandInOut = trigger('expandInOut', [
    state('true', style({height: '0', opacity: '0'})),
    state('false', style({height: AUTO_STYLE, opacity: '1'})),
    transition('false => true', [style({overflow: 'hidden'}), animate('500ms ease-in', style({height: '0'}))]),
  ]);
}