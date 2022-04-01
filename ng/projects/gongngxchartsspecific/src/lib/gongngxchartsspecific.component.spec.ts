import { ComponentFixture, TestBed } from '@angular/core/testing';

import { GongngxchartsspecificComponent } from './gongngxchartsspecific.component';

describe('GongngxchartsspecificComponent', () => {
  let component: GongngxchartsspecificComponent;
  let fixture: ComponentFixture<GongngxchartsspecificComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ GongngxchartsspecificComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(GongngxchartsspecificComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
