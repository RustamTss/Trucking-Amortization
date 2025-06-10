export interface User {
  id: string;
  email: string;
  name: string;
  companies: string[];
  created_at: string;
}

export interface Company {
  id: string;
  name: string;
  owner_id: string;
  created_at: string;
}

export interface LoanInfo {
  loan_amount: number;
  interest_rate: number;
  loan_term: number;
  start_date: string;
  monthly_payment: number;
}

export interface Asset {
  id: string;
  company_id: string;
  type: 'truck' | 'trailer';
  make: string;
  model: string;
  year: number;
  vin: string;
  purchase_date: string;
  purchase_price: number;
  loan_info?: LoanInfo;
  created_at: string;
}

export interface AmortizationEntry {
  month: number;
  payment: number;
  principal: number;
  interest: number;
  balance: number;
  cumulative_interest: number;
  date: string;
}

export interface AmortizationSchedule {
  asset_id: string;
  total_amount: number;
  monthly_payment: number;
  total_interest: number;
  entries: AmortizationEntry[];
}

export interface DepreciationEntry {
  year: number;
  depreciation_amount: number;
  accumulated_depreciation: number;
  book_value: number;
  date: string;
}

export interface DepreciationSchedule {
  asset_id: string;
  initial_value: number;
  useful_life: number;
  annual_depreciation: number;
  entries: DepreciationEntry[];
}

export interface BusinessDebtSummary {
  company_id: string;
  total_loan_amount: number;
  total_balance: number;
  monthly_payment: number;
  assets_count: number;
}

export interface BusinessDebtDetail {
  asset_id: string;
  asset_type: string;
  make: string;
  model: string;
  year: number;
  loan_amount: number;
  current_balance: number;
  monthly_payment: number;
  interest_rate: number;
  remaining_months: number;
}

export interface BusinessDebtSchedule {
  summary: BusinessDebtSummary;
  details: BusinessDebtDetail[];
}

export interface AuthResponse {
  user: User;
  token: string;
}

export interface LoginRequest {
  email: string;
  password: string;
}

export interface RegisterRequest {
  email: string;
  password: string;
  name: string;
}

export interface AssetRequest {
  company_id: string;
  type: 'truck' | 'trailer';
  make: string;
  model: string;
  year: number;
  vin: string;
  purchase_date: string;
  purchase_price: number;
  loan_info?: LoanInfo;
}

export interface CompanyRequest {
  name: string;
} 