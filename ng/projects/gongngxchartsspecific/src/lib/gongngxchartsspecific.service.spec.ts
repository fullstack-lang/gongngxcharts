import { TestBed } from '@angular/core/testing';

import { GongngxchartsspecificService } from './gongngxchartsspecific.service';

describe('GongngxchartsspecificService', () => {
  let service: GongngxchartsspecificService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(GongngxchartsspecificService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
