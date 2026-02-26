import client from './client';

export interface StatsCounts {
  orderCount: number;
  materialCount: number;
  specCount: number;
  vendorCount: number;
}

export interface YearCount {
  year: number;
  count: number;
}

export interface MonthCount {
  month: number;
  count: number;
}

export interface NameCount {
  name: string;
  count: number;
}

export interface StatsResponse {
  counts: StatsCounts;
  ordersByYear: YearCount[];
  ordersByMonthCur: MonthCount[];
  ordersByMonthLast: MonthCount[];
  ordersByMaterial: NameCount[];
  ordersBySpec: NameCount[];
}

export function fetchStats() {
  return client.get<StatsResponse>('/stats');
}
