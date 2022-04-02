import { NgModule } from '@angular/core';
import { GongngxchartsspecificComponent } from './gongngxchartsspecific.component';
import { GongngxchartsChartComponent } from './gongngxcharts-chart/gongngxcharts-chart.component';

import { NgChartsModule } from 'ng2-charts';

@NgModule({
  declarations: [
    GongngxchartsspecificComponent,
    GongngxchartsChartComponent
  ],
  imports: [
    NgChartsModule,
  ],
  exports: [
    GongngxchartsspecificComponent,
    GongngxchartsChartComponent
  ]
})
export class GongngxchartsspecificModule { }
