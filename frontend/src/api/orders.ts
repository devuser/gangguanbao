import client from './client';

export interface Order {
  id: number;
  orderNo: string;
  materialId?: number;
  specId?: number;
  vendorId?: number;
  orderDate?: string;
  quantity: number;
  unitPrice?: number;
  amount?: number;
  remark: string;
  materialName: string;
  specName: string;
  vendorName: string;
}

export interface OrderQuery {
  materialId?: number;
  specId?: number;
  vendorId?: number;
  orderNo?: string;
  dateFrom?: string;
  dateTo?: string;
  page?: number;
  pageSize?: number;
  sortBy?: string;
  sortOrder?: 'asc' | 'desc';
}

export interface OrderListResponse {
  list: Order[];
  total: number;
}

export function fetchOrders(params: OrderQuery) {
  return client.get<OrderListResponse>('/orders', { params });
}

export function fetchOrder(id: number) {
  return client.get<Order>(`/orders/${id}`);
}

export function createOrder(data: Partial<Order>) {
  return client.post<{ id: number }>('/orders', data);
}

export function updateOrder(id: number, data: Partial<Order>) {
  return client.put(`/orders/${id}`, data);
}

export function deleteOrder(id: number) {
  return client.delete(`/orders/${id}`);
}

export function exportOrders(params: OrderQuery) {
  return client.get('/orders/export', { params, responseType: 'blob' });
}
