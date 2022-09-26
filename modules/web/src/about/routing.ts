import { NgModule } from "@angular/core";
import { Route, RouterModule } from "@angular/router";
import {BREADCRUMBS} from '../index.messages';
import { AboutComponent } from "./component";

export const ABOUT_ROUTE: Route = {
  path: '',
  component: AboutComponent,
  data: {
    breadcrumb: BREADCRUMBS.Create,
  },
};

export const ACTIONBAR: Route = {
  path: '',
  component: ActionbarComponent,
  outlet: 'actionbar'
}

@NgModule({
  imports: [RouterModule.forChild([ABOUT_ROUTE, ACTIONBAR])],
  exports: [RouterModule]
})
export class AboutRoutingModule {}