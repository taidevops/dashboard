import { InjectionToken } from "@angular/core";
import { IMessage } from "@api/root.ui";

export const MESSAGES_DI_TOKEN = new InjectionToken<IMessage>('kd.messages');
