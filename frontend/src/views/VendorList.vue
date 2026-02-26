<template>
  <PageLayout title="厂家管理" desc="维护厂家/供应商基础数据，供订单查询时选择">
    <template #actions>
      <button class="btn btn-primary" @click="showCreate = true">新增</button>
      <button class="btn" @click="doExport">导出Excel</button>
    </template>
    <div class="vendor-list">
      <div class="card search-card">
        <div class="search-row">
          <div class="field">
            <label>厂家名称</label>
            <input v-model="searchName" type="text" class="input" placeholder="模糊搜索" />
          </div>
          <div class="field">
            <label>编码</label>
            <input v-model="searchCode" type="text" class="input" placeholder="模糊搜索" />
          </div>
          <div class="actions">
            <button class="btn btn-primary" @click="search">查询</button>
            <button class="btn" @click="reset">重置</button>
          </div>
        </div>
      </div>

      <div v-if="errorMsg" class="error-msg">{{ errorMsg }}</div>

      <div class="card table-card">
        <div class="table-wrap">
          <table class="data-table">
            <thead>
              <tr>
                <th
                  v-for="col in columns"
                  :key="col.key"
                  :class="{ sortable: col.sortable }"
                  @click="col.sortable && sort(col.key)"
                >
                  {{ col.label }}
                  <span v-if="col.sortable && sortBy === col.key" class="sort-icon">
                    {{ sortOrder === 'asc' ? '↑' : '↓' }}
                  </span>
                </th>
                <th class="th-actions">操作</th>
              </tr>
            </thead>
            <tbody>
              <tr v-if="loading">
                <td :colspan="columns.length + 1" class="loading-cell">加载中...</td>
              </tr>
              <tr v-else-if="pagedList.length === 0">
                <td :colspan="columns.length + 1" class="empty-cell">暂无数据</td>
              </tr>
              <tr v-else v-for="v in pagedList" :key="v.id">
                <td>{{ v.id }}</td>
                <td>{{ v.name || '-' }}</td>
                <td>{{ v.code || '-' }}</td>
                <td>{{ v.contact || '-' }}</td>
                <td>{{ v.phone || '-' }}</td>
                <td>{{ formatDate(v.createdAt) }}</td>
                <td>
                  <button class="btn-link" @click="openDetail(v)">详情</button>
                  <span class="btn-sep">|</span>
                  <button class="btn-link" @click="openEdit(v)">编辑</button>
                  <span class="btn-sep">|</span>
                  <button class="btn-link btn-danger" @click="confirmDelete(v)">删除</button>
                </td>
              </tr>
            </tbody>
          </table>
        </div>

        <Pagination
          :total="filteredList.length"
          :page="page"
          :page-size="pageSize"
          :preset-sizes="[10, 20, 30, 40, 50, 100]"
          :allow-custom="true"
          @update:page="page = $event"
          @update:page-size="
            pageSize = $event;
            page = 1;
          "
        />
      </div>
    </div>

    <Modal
      :model-value="showCreate"
      @update:model-value="
        (v) => {
          showCreate = v;
          if (!v) resetForm();
        }
      "
      title="新增厂家"
    >
      <div class="form-row">
        <div class="field">
          <label>厂家名称 <span class="required">*</span></label>
          <input v-model="form.name" type="text" class="input" placeholder="请输入" />
        </div>
        <div class="field">
          <label>编码</label>
          <input v-model="form.code" type="text" class="input" placeholder="可选" />
        </div>
        <div class="field">
          <label>联系人</label>
          <input v-model="form.contact" type="text" class="input" placeholder="可选" />
        </div>
        <div class="field">
          <label>电话</label>
          <input v-model="form.phone" type="text" class="input" placeholder="可选" />
        </div>
      </div>
      <template #footer>
        <button class="btn" @click="showCreate = false">取消</button>
        <button class="btn btn-primary" @click="doCreate">确定</button>
      </template>
    </Modal>

    <Modal
      :model-value="showDetail"
      @update:model-value="
        (v) => {
          showDetail = v;
          if (!v) detailTarget = null;
        }
      "
      title="厂家详情"
    >
      <div v-if="detailTarget" class="detail-grid">
        <div class="detail-item">
          <span class="detail-label">ID</span><span>{{ detailTarget.id }}</span>
        </div>
        <div class="detail-item">
          <span class="detail-label">厂家名称</span><span>{{ detailTarget.name || '-' }}</span>
        </div>
        <div class="detail-item">
          <span class="detail-label">编码</span><span>{{ detailTarget.code || '-' }}</span>
        </div>
        <div class="detail-item">
          <span class="detail-label">联系人</span><span>{{ detailTarget.contact || '-' }}</span>
        </div>
        <div class="detail-item">
          <span class="detail-label">电话</span><span>{{ detailTarget.phone || '-' }}</span>
        </div>
        <div class="detail-item">
          <span class="detail-label">创建时间</span
          ><span>{{ formatDate(detailTarget.createdAt) }}</span>
        </div>
      </div>
      <template #footer>
        <button class="btn btn-primary" @click="showDetail = false">关闭</button>
      </template>
    </Modal>

    <Modal
      :model-value="showEdit"
      @update:model-value="
        (v) => {
          showEdit = v;
          if (!v) editTarget = null;
        }
      "
      title="编辑厂家"
    >
      <div v-if="editTarget" class="form-row">
        <div class="field">
          <label>厂家名称 <span class="required">*</span></label>
          <input v-model="editForm.name" type="text" class="input" placeholder="请输入" />
        </div>
        <div class="field">
          <label>编码</label>
          <input v-model="editForm.code" type="text" class="input" placeholder="可选" />
        </div>
        <div class="field">
          <label>联系人</label>
          <input v-model="editForm.contact" type="text" class="input" placeholder="可选" />
        </div>
        <div class="field">
          <label>电话</label>
          <input v-model="editForm.phone" type="text" class="input" placeholder="可选" />
        </div>
      </div>
      <template #footer>
        <button class="btn" @click="showEdit = false">取消</button>
        <button class="btn btn-primary" @click="doUpdate">保存</button>
      </template>
    </Modal>

    <Modal
      :model-value="showDelete"
      @update:model-value="
        (v) => {
          showDelete = v;
          if (!v) deleteTarget = null;
        }
      "
      title="确认删除"
    >
      <p>确定要删除厂家「{{ deleteTarget?.name }}」吗？</p>
      <template #footer>
        <button class="btn" @click="showDelete = false">取消</button>
        <button class="btn btn-danger" @click="doDelete">确定删除</button>
      </template>
    </Modal>
  </PageLayout>
</template>

<script setup lang="ts">
import { ref, reactive, computed, onMounted } from 'vue';
import {
  fetchVendors,
  createVendor,
  updateVendor,
  deleteVendor,
  exportVendors,
  type Vendor,
} from '@/api/meta';
import { downloadBlob, getExportFilename } from '@/utils/export';
import PageLayout from '@/components/PageLayout.vue';
import Pagination from '@/components/Pagination.vue';
import Modal from '@/components/Modal.vue';

const vendors = ref<Vendor[]>([]);
const loading = ref(false);
const errorMsg = ref('');
const searchName = ref('');
const searchCode = ref('');
const page = ref(1);
const pageSize = ref(10);
const sortBy = ref('name');
const sortOrder = ref<'asc' | 'desc'>('asc');
const showCreate = ref(false);
const showDetail = ref(false);
const showEdit = ref(false);
const showDelete = ref(false);
const detailTarget = ref<Vendor | null>(null);
const editTarget = ref<Vendor | null>(null);
const deleteTarget = ref<Vendor | null>(null);

const form = reactive({ name: '', code: '', contact: '', phone: '' });
const editForm = reactive({ name: '', code: '', contact: '', phone: '' });

function resetForm() {
  form.name = '';
  form.code = '';
  form.contact = '';
  form.phone = '';
}

const columns = [
  { key: 'id', label: 'ID', sortable: true },
  { key: 'name', label: '厂家名称', sortable: true },
  { key: 'code', label: '编码', sortable: true },
  { key: 'contact', label: '联系人', sortable: true },
  { key: 'phone', label: '电话', sortable: true },
  { key: 'createdAt', label: '创建时间', sortable: true },
];

const filteredList = computed(() => {
  let list = vendors.value;
  const name = searchName.value.trim().toLowerCase();
  const code = searchCode.value.trim().toLowerCase();
  if (name) list = list.filter((v) => (v.name || '').toLowerCase().includes(name));
  if (code) list = list.filter((v) => (v.code || '').toLowerCase().includes(code));
  return list;
});

const sortedList = computed(() => {
  const list = [...filteredList.value];
  const key = sortBy.value;
  const ord = sortOrder.value === 'asc' ? 1 : -1;
  list.sort((a, b) => {
    const va = (a as Record<string, unknown>)[key];
    const vb = (b as Record<string, unknown>)[key];
    if (va == null && vb == null) return 0;
    if (va == null) return ord;
    if (vb == null) return -ord;
    if (key === 'id' && typeof va === 'number' && typeof vb === 'number') {
      return (va - vb) * ord;
    }
    return String(va).localeCompare(String(vb)) * ord;
  });
  return list;
});

const pagedList = computed(() => {
  const start = (page.value - 1) * pageSize.value;
  return sortedList.value.slice(start, start + pageSize.value);
});

function formatDate(d: string | undefined) {
  if (!d) return '-';
  try {
    return new Date(d).toLocaleDateString('zh-CN');
  } catch {
    return d;
  }
}

function search() {
  page.value = 1;
}

function reset() {
  searchName.value = '';
  searchCode.value = '';
  page.value = 1;
}

function sort(col: string) {
  if (sortBy.value === col) {
    sortOrder.value = sortOrder.value === 'asc' ? 'desc' : 'asc';
  } else {
    sortBy.value = col;
    sortOrder.value = 'asc';
  }
  page.value = 1;
}

async function load() {
  loading.value = true;
  errorMsg.value = '';
  try {
    const res = await fetchVendors();
    vendors.value = res.data.list ?? [];
  } catch (e: unknown) {
    errorMsg.value = e instanceof Error ? e.message : '加载失败';
  } finally {
    loading.value = false;
  }
}

async function doCreate() {
  if (!form.name.trim()) {
    errorMsg.value = '厂家名称为必填';
    return;
  }
  try {
    await createVendor({
      name: form.name.trim(),
      code: form.code.trim(),
      contact: form.contact.trim(),
      phone: form.phone.trim(),
    });
    showCreate.value = false;
    resetForm();
    load();
  } catch (e: unknown) {
    errorMsg.value = e instanceof Error ? e.message : '新增失败';
  }
}

function openDetail(v: Vendor) {
  detailTarget.value = v;
  showDetail.value = true;
}

function openEdit(v: Vendor) {
  editTarget.value = v;
  editForm.name = v.name || '';
  editForm.code = v.code || '';
  editForm.contact = v.contact || '';
  editForm.phone = v.phone || '';
  showEdit.value = true;
}

async function doUpdate() {
  if (!editTarget.value) return;
  if (!editForm.name.trim()) {
    errorMsg.value = '厂家名称为必填';
    return;
  }
  try {
    await updateVendor(editTarget.value.id, {
      name: editForm.name.trim(),
      code: editForm.code.trim(),
      contact: editForm.contact.trim(),
      phone: editForm.phone.trim(),
    });
    showEdit.value = false;
    editTarget.value = null;
    load();
  } catch (e: unknown) {
    errorMsg.value = e instanceof Error ? e.message : '更新失败';
  }
}

function confirmDelete(v: Vendor) {
  deleteTarget.value = v;
  showDelete.value = true;
}

async function doDelete() {
  if (!deleteTarget.value) return;
  try {
    await deleteVendor(deleteTarget.value.id);
    showDelete.value = false;
    deleteTarget.value = null;
    load();
  } catch (e: unknown) {
    errorMsg.value = e instanceof Error ? e.message : '删除失败';
  }
}

async function doExport() {
  try {
    const res = await exportVendors();
    const filename = getExportFilename(res.headers['content-disposition'] as string);
    downloadBlob(new Blob([res.data]), filename);
  } catch (e: unknown) {
    errorMsg.value = e instanceof Error ? e.message : '导出失败';
  }
}

onMounted(load);
</script>

<style scoped>
.vendor-list {
  display: flex;
  flex-direction: column;
  gap: 1rem;
  flex: 1;
  min-height: 0;
}

.search-card {
  padding: 1.25rem;
  flex-shrink: 0;
}

.search-row {
  display: flex;
  flex-wrap: wrap;
  gap: 1rem;
  align-items: flex-end;
}

.field {
  display: flex;
  flex-direction: column;
  gap: 0.25rem;
}

.field label {
  font-size: 0.85rem;
  color: var(--color-text-muted);
}

.required {
  color: var(--color-error);
}

.input {
  padding: 0.5rem 0.75rem;
  border: 1px solid var(--color-border);
  border-radius: var(--radius-sm);
  min-width: 140px;
}

.actions {
  display: flex;
  gap: 0.5rem;
}

.form-row {
  display: flex;
  flex-wrap: wrap;
  gap: 1rem;
}

.form-row .field {
  min-width: 200px;
}

.table-card {
  flex: 1;
  display: flex;
  flex-direction: column;
  min-height: 0;
  overflow: hidden;
}

.table-wrap {
  flex: 1;
  overflow: auto;
}

.data-table th,
.data-table td {
  padding: 0.75rem 1rem;
  text-align: left;
  border-bottom: 1px solid var(--color-border);
}

.data-table th {
  background: var(--color-bg);
  font-weight: 500;
  white-space: nowrap;
}

.th-actions {
  width: 80px;
}

.data-table th.sortable {
  cursor: pointer;
  user-select: none;
}

.data-table th.sortable:hover {
  color: var(--color-accent);
}

.sort-icon {
  margin-left: 0.25rem;
  font-size: 0.8em;
}

.btn-link {
  background: none;
  border: none;
  color: var(--color-accent);
  cursor: pointer;
  font-size: 0.9rem;
}

.btn-link:hover {
  text-decoration: underline;
}

.btn-sep {
  margin: 0 0.35rem;
  color: var(--color-border);
}

.btn-danger {
  color: var(--color-error);
}

.detail-grid {
  display: grid;
  gap: 0.75rem 1.5rem;
  grid-template-columns: auto 1fr;
}

.detail-item {
  display: contents;
}

.detail-label {
  color: var(--color-text-muted);
  font-size: 0.9rem;
}

.loading-cell,
.empty-cell {
  text-align: center;
  padding: 2rem;
  color: var(--color-text-muted);
}

.error-msg {
  padding: 0.75rem 1rem;
  background: #ffebe9;
  color: var(--color-error);
  border-radius: var(--radius-sm);
  font-size: 0.9rem;
}
</style>
