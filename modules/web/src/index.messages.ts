import { InjectionToken } from "@angular/core";
import { IBreadcrumbMessage, IBreadcrumbMessageKey, IMessage } from "@api/root.ui";

export const MESSAGES_DI_TOKEN = new InjectionToken<IMessage>('kd.messages');

export const BREADCRUMBS: IBreadcrumbMessage = {
  [IBreadcrumbMessageKey.Logs]: $localize`Logs`,
  [IBreadcrumbMessageKey.Error]: $localize`Error`,
  [IBreadcrumbMessageKey.Create]: $localize`Create`,
}