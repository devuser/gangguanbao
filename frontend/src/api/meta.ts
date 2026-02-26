import client from './client';

export interface Material {
  id: number;
  name: string;
  code: string;
  createdAt?: string;
}

export interface Spec {
  id: number;
  name: string;
  materialId?: number;
  materialName?: string;
  unit: string;
  createdAt?: string;
}

export interface Vendor {
  id: number;
  name: string;
  code: string;
  contact: string;
  phone: string;
  createdAt?: string;
}

export function fetchMaterials() {
  return client.get<{ list: Material[] }>('/materials');
}

export function fetchMaterial(id: number) {
  return client.get<Material>(`/materials/${id}`);
}

export function createMaterial(data: Partial<Material>) {
  return client.post<{ id: number }>('/materials', data);
}

export function updateMaterial(id: number, data: Partial<Material>) {
  return client.put(`/materials/${id}`, data);
}

export function deleteMaterial(id: number) {
  return client.delete(`/materials/${id}`);
}

export function exportMaterials() {
  return client.get('/materials/export', { responseType: 'blob' });
}

export function fetchSpecs() {
  return client.get<{ list: Spec[] }>('/specs');
}

export function fetchSpec(id: number) {
  return client.get<Spec>(`/specs/${id}`);
}

export function createSpec(data: Partial<Spec>) {
  return client.post<{ id: number }>('/specs', data);
}

export function updateSpec(id: number, data: Partial<Spec>) {
  return client.put(`/specs/${id}`, data);
}

export function deleteSpec(id: number) {
  return client.delete(`/specs/${id}`);
}

export function exportSpecs() {
  return client.get('/specs/export', { responseType: 'blob' });
}

export function fetchVendors() {
  return client.get<{ list: Vendor[] }>('/vendors');
}

export function fetchVendor(id: number) {
  return client.get<Vendor>(`/vendors/${id}`);
}

export function createVendor(data: Partial<Vendor>) {
  return client.post<{ id: number }>('/vendors', data);
}

export function updateVendor(id: number, data: Partial<Vendor>) {
  return client.put(`/vendors/${id}`, data);
}

export function deleteVendor(id: number) {
  return client.delete(`/vendors/${id}`);
}

export function exportVendors() {
  return client.get('/vendors/export', { responseType: 'blob' });
}
