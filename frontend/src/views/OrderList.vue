<template>
  <PageLayout title="订单查询" desc="按材质、规格、厂家等条件查询往年订单数据">
    <template #actions>
      <button class="btn btn-primary" @click="showCreate = true">新增</button>
      <button class="btn" @click="doExport">导出Excel</button>
    </template>
    <div class="order-list">
      <div class="card search-card">
        <div class="search-row">
          <div class="field">
            <label>材质</label>
            <select v-model="query.materialId" class="input">
              <option :value="undefined">全部</option>
              <option v-for="m in materials" :key="m.id" :value="m.id">{{ m.name }}</option>
            </select>
          </div>
          <div class="field">
            <label>规格</label>
            <select v-model="query.specId" class="input">
              <option :value="undefined">全部</option>
              <option v-for="s in specs" :key="s.id" :value="s.id">{{ s.name }}</option>
            </select>
          </div>
          <div class="field">
            <label>厂家</label>
            <select v-model="query.vendorId" class="input">
              <option :value="undefined">全部</option>
              <option v-for="v in vendors" :key="v.id" :value="v.id">{{ v.name }}</option>
            </select>
          </div>
          <div class="field">
            <label>订单号</label>
            <input v-model="query.orderNo" type="text" class="input" placeholder="模糊搜索" />
          </div>
          <div class="field">
            <label>日期起</label>
            <input v-model="query.dateFrom" type="date" class="input" />
          </div>
          <div class="field">
            <label>日期止</label>
            <input v-model="query.dateTo" type="date" class="input" />
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
              <tr v-else-if="orders.length === 0">
                <td :colspan="columns.length + 1" class="empty-cell">暂无数据</td>
              </tr>
              <tr v-else v-for="o in orders" :key="o.id">
                <td>{{ o.orderNo || '-' }}</td>
                <td>{{ formatDate(o.orderDate) }}</td>
                <td>{{ o.materialName || '-' }}</td>
                <td>{{ o.specName || '-' }}</td>
                <td>{{ o.vendorName || '-' }}</td>
                <td>{{ formatNum(o.quantity) }}</td>
                <td>{{ formatMoney(o.unitPrice) }}</td>
                <td>{{ formatMoney(o.amount) }}</td>
                <td>{{ o.remark || '-' }}</td>
                <td>
                  <button class="btn-link" @click="openDetail(o)">详情</button>
                  <span class="btn-sep">|</span>
                  <button class="btn-link" @click="openEdit(o)">编辑</button>
                  <span class="btn-sep">|</span>
                  <button class="btn-link btn-danger" @click="confirmDelete(o)">删除</button>
                </td>
              </tr>
            </tbody>
          </table>
        </div>

        <Pagination
          :total="total"
          :page="query.page"
          :page-size="query.pageSize"
          :preset-sizes="[10, 20, 30, 40, 50, 100]"
          :allow-custom="true"
          @update:page="
            query.page = $event;
            load();
          "
          @update:page-size="
            query.pageSize = $event;
            load();
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
      title="新增订单"
    >
      <div class="form-row">
        <div class="field">
          <label>订单号 <span class="required">*</span></label>
          <input v-model="form.orderNo" type="text" class="input" placeholder="请输入" />
        </div>
        <div class="field">
          <label>材质</label>
          <select v-model="form.materialId" class="input">
            <option :value="undefined">请选择</option>
            <option v-for="m in materials" :key="m.id" :value="m.id">{{ m.name }}</option>
          </select>
        </div>
        <div class="field">
          <label>规格</label>
          <select v-model="form.specId" class="input">
            <option :value="undefined">请选择</option>
            <option v-for="s in specs" :key="s.id" :value="s.id">{{ s.name }}</option>
          </select>
        </div>
        <div class="field">
          <label>厂家</label>
          <select v-model="form.vendorId" class="input">
            <option :value="undefined">请选择</option>
            <option v-for="v in vendors" :key="v.id" :value="v.id">{{ v.name }}</option>
          </select>
        </div>
        <div class="field">
          <label>订单日期</label>
          <input v-model="form.orderDate" type="date" class="input" />
        </div>
        <div class="field">
          <label>数量</label>
          <input
            v-model.number="form.quantity"
            type="number"
            step="0.01"
            class="input"
            placeholder="0"
          />
        </div>
        <div class="field">
          <label>单价</label>
          <input
            v-model.number="form.unitPrice"
            type="number"
            step="0.01"
            class="input"
            placeholder="可选"
          />
        </div>
        <div class="field">
          <label>金额</label>
          <input
            v-model.number="form.amount"
            type="number"
            step="0.01"
            class="input"
            placeholder="可选"
          />
        </div>
        <div class="field full-width">
          <label>备注</label>
          <input v-model="form.remark" type="text" class="input" placeholder="可选" />
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
      title="订单详情"
    >
      <div v-if="detailTarget" class="detail-grid">
        <div class="detail-item">
          <span class="detail-label">ID</span><span>{{ detailTarget.id ?? '-' }}</span>
        </div>
        <div class="detail-item">
          <span class="detail-label">订单号</span><span>{{ detailTarget.orderNo || '-' }}</span>
        </div>
        <div class="detail-item">
          <span class="detail-label">订单日期</span
          ><span>{{ formatDate(detailTarget.orderDate) }}</span>
        </div>
        <div class="detail-item">
          <span class="detail-label">材质</span><span>{{ detailTarget.materialName || '-' }}</span>
        </div>
        <div class="detail-item">
          <span class="detail-label">规格</span><span>{{ detailTarget.specName || '-' }}</span>
        </div>
        <div class="detail-item">
          <span class="detail-label">厂家</span><span>{{ detailTarget.vendorName || '-' }}</span>
        </div>
        <div class="detail-item">
          <span class="detail-label">数量</span><span>{{ formatNum(detailTarget.quantity) }}</span>
        </div>
        <div class="detail-item">
          <span class="detail-label">单价</span
          ><span>{{ formatMoney(detailTarget.unitPrice) }}</span>
        </div>
        <div class="detail-item">
          <span class="detail-label">金额</span><span>{{ formatMoney(detailTarget.amount) }}</span>
        </div>
        <div class="detail-item">
          <span class="detail-label">备注</span><span>{{ detailTarget.remark || '-' }}</span>
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
      title="编辑订单"
    >
      <div v-if="editTarget" class="form-row">
        <div class="field">
          <label>订单号 <span class="required">*</span></label>
          <input v-model="editForm.orderNo" type="text" class="input" placeholder="请输入" />
        </div>
        <div class="field">
          <label>材质</label>
          <select v-model="editForm.materialId" class="input">
            <option :value="undefined">请选择</option>
            <option v-for="m in materials" :key="m.id" :value="m.id">{{ m.name }}</option>
          </select>
        </div>
        <div class="field">
          <label>规格</label>
          <select v-model="editForm.specId" class="input">
            <option :value="undefined">请选择</option>
            <option v-for="s in specs" :key="s.id" :value="s.id">{{ s.name }}</option>
          </select>
        </div>
        <div class="field">
          <label>厂家</label>
          <select v-model="editForm.vendorId" class="input">
            <option :value="undefined">请选择</option>
            <option v-for="v in vendors" :key="v.id" :value="v.id">{{ v.name }}</option>
          </select>
        </div>
        <div class="field">
          <label>订单日期</label>
          <input v-model="editForm.orderDate" type="date" class="input" />
        </div>
        <div class="field">
          <label>数量</label>
          <input
            v-model.number="editForm.quantity"
            type="number"
            step="0.01"
            class="input"
            placeholder="0"
          />
        </div>
        <div class="field">
          <label>单价</label>
          <input
            v-model.number="editForm.unitPrice"
            type="number"
            step="0.01"
            class="input"
            placeholder="可选"
          />
        </div>
        <div class="field">
          <label>金额</label>
          <input
            v-model.number="editForm.amount"
            type="number"
            step="0.01"
            class="input"
            placeholder="可选"
          />
        </div>
        <div class="field full-width">
          <label>备注</label>
          <input v-model="editForm.remark" type="text" class="input" placeholder="可选" />
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
      <p>确定要删除订单「{{ deleteTarget?.orderNo }}」吗？</p>
      <template #footer>
        <button class="btn" @click="showDelete = false">取消</button>
        <button class="btn btn-danger" @click="doDelete">确定删除</button>
      </template>
    </Modal>
  </PageLayout>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted } from 'vue';
import {
  fetchOrders,
  createOrder,
  updateOrder,
  deleteOrder,
  exportOrders,
  type Order,
  type OrderQuery,
} from '@/api/orders';
import {
  fetchMaterials,
  fetchSpecs,
  fetchVendors,
  type Material,
  type Spec,
  type Vendor,
} from '@/api/meta';
import { downloadBlob, getExportFilename } from '@/utils/export';
import PageLayout from '@/components/PageLayout.vue';
import Pagination from '@/components/Pagination.vue';
import Modal from '@/components/Modal.vue';

const materials = ref<Material[]>([]);
const specs = ref<Spec[]>([]);
const vendors = ref<Vendor[]>([]);
const orders = ref<Order[]>([]);
const total = ref(0);
const loading = ref(false);
const errorMsg = ref('');

const query = reactive<OrderQuery & { page: number; pageSize: number }>({
  materialId: undefined,
  specId: undefined,
  vendorId: undefined,
  orderNo: '',
  dateFrom: '',
  dateTo: '',
  page: 1,
  pageSize: 10,
  sortBy: 'orderDate',
  sortOrder: 'desc',
});

const sortBy = ref('orderDate');
const sortOrder = ref<'asc' | 'desc'>('desc');
const showCreate = ref(false);
const showDetail = ref(false);
const showEdit = ref(false);
const showDelete = ref(false);
const detailTarget = ref<Order | null>(null);
const editTarget = ref<Order | null>(null);
const deleteTarget = ref<Order | null>(null);

const form = reactive({
  orderNo: '',
  materialId: undefined as number | undefined,
  specId: undefined as number | undefined,
  vendorId: undefined as number | undefined,
  orderDate: '',
  quantity: 0,
  unitPrice: undefined as number | undefined,
  amount: undefined as number | undefined,
  remark: '',
});

const editForm = reactive({
  orderNo: '',
  materialId: undefined as number | undefined,
  specId: undefined as number | undefined,
  vendorId: undefined as number | undefined,
  orderDate: '',
  quantity: 0,
  unitPrice: undefined as number | undefined,
  amount: undefined as number | undefined,
  remark: '',
});

function resetForm() {
  form.orderNo = '';
  form.materialId = undefined;
  form.specId = undefined;
  form.vendorId = undefined;
  form.orderDate = '';
  form.quantity = 0;
  form.unitPrice = undefined;
  form.amount = undefined;
  form.remark = '';
}

const columns = [
  { key: 'orderNo', label: '订单号', sortable: true },
  { key: 'orderDate', label: '订单日期', sortable: true },
  { key: 'materialName', label: '材质', sortable: true },
  { key: 'specName', label: '规格', sortable: true },
  { key: 'vendorName', label: '厂家', sortable: true },
  { key: 'quantity', label: '数量', sortable: true },
  { key: 'unitPrice', label: '单价', sortable: true },
  { key: 'amount', label: '金额', sortable: true },
  { key: 'remark', label: '备注', sortable: false },
];

function formatDate(d: string | undefined): string {
  if (!d) return '-';
  try {
    const date = new Date(d);
    return isNaN(date.getTime()) ? String(d) : date.toLocaleDateString('zh-CN');
  } catch {
    return String(d ?? '-');
  }
}

/** 转为 input type="date" 所需的 YYYY-MM-DD 格式 */
function toDateInputValue(d: string | undefined): string {
  if (!d) return '';
  const m = d.match(/^(\d{4})-(\d{2})-(\d{2})/);
  if (m) return m[0];
  try {
    const date = new Date(d);
    if (!isNaN(date.getTime())) {
      return date.toISOString().slice(0, 10);
    }
  } catch {
    /* ignore */
  }
  return '';
}

function formatNum(val: number | undefined | null): string {
  if (val == null || (typeof val === 'number' && isNaN(val))) return '-';
  return String(val);
}

function formatMoney(val: number | undefined | null): string {
  if (val == null || (typeof val === 'number' && isNaN(val))) return '-';
  return Number(val).toFixed(2);
}

async function loadMeta() {
  try {
    const [mRes, sRes, vRes] = await Promise.all([fetchMaterials(), fetchSpecs(), fetchVendors()]);
    materials.value = mRes.data.list ?? [];
    specs.value = sRes.data.list ?? [];
    vendors.value = vRes.data.list ?? [];
  } catch (e: unknown) {
    errorMsg.value = e instanceof Error ? e.message : '加载基础数据失败';
  }
}

async function load() {
  loading.value = true;
  errorMsg.value = '';
  try {
    const res = await fetchOrders({
      ...query,
      sortBy: sortBy.value,
      sortOrder: sortOrder.value,
    });
    orders.value = res.data.list ?? [];
    total.value = res.data.total ?? 0;
  } catch (e: unknown) {
    errorMsg.value = e instanceof Error ? e.message : '查询失败，请检查网络或后端服务';
  } finally {
    loading.value = false;
  }
}

function search() {
  query.page = 1;
  load();
}

function reset() {
  query.materialId = undefined;
  query.specId = undefined;
  query.vendorId = undefined;
  query.orderNo = '';
  query.dateFrom = '';
  query.dateTo = '';
  query.page = 1;
  load();
}

function sort(col: string) {
  if (sortBy.value === col) {
    sortOrder.value = sortOrder.value === 'asc' ? 'desc' : 'asc';
  } else {
    sortBy.value = col;
    sortOrder.value = 'asc';
  }
  query.page = 1;
  load();
}

async function doCreate() {
  if (!form.orderNo.trim()) {
    errorMsg.value = '订单号为必填';
    return;
  }
  try {
    const payload: Record<string, unknown> = {
      orderNo: form.orderNo.trim(),
      quantity: form.quantity ?? 0,
      remark: form.remark.trim(),
    };
    if (form.materialId != null) payload.materialId = form.materialId;
    if (form.specId != null) payload.specId = form.specId;
    if (form.vendorId != null) payload.vendorId = form.vendorId;
    if (form.orderDate) payload.orderDate = form.orderDate;
    if (form.unitPrice != null) payload.unitPrice = form.unitPrice;
    if (form.amount != null) payload.amount = form.amount;
    await createOrder(payload as Partial<Order>);
    showCreate.value = false;
    resetForm();
    load();
  } catch (e: unknown) {
    errorMsg.value = e instanceof Error ? e.message : '新增失败';
  }
}

function openDetail(o: Order) {
  detailTarget.value = o;
  showDetail.value = true;
}

function openEdit(o: Order) {
  editTarget.value = o;
  editForm.orderNo = o.orderNo || '';
  editForm.materialId = o.materialId;
  editForm.specId = o.specId;
  editForm.vendorId = o.vendorId;
  editForm.orderDate = toDateInputValue(o.orderDate);
  editForm.quantity = o.quantity ?? 0;
  editForm.unitPrice = o.unitPrice;
  editForm.amount = o.amount;
  editForm.remark = o.remark || '';
  showEdit.value = true;
}

async function doUpdate() {
  if (!editTarget.value) return;
  if (!editForm.orderNo.trim()) {
    errorMsg.value = '订单号为必填';
    return;
  }
  try {
    const payload: Record<string, unknown> = {
      orderNo: editForm.orderNo.trim(),
      quantity: editForm.quantity ?? 0,
      remark: editForm.remark.trim(),
    };
    if (editForm.materialId != null) payload.materialId = editForm.materialId;
    if (editForm.specId != null) payload.specId = editForm.specId;
    if (editForm.vendorId != null) payload.vendorId = editForm.vendorId;
    if (editForm.orderDate) payload.orderDate = editForm.orderDate;
    if (editForm.unitPrice != null) payload.unitPrice = editForm.unitPrice;
    if (editForm.amount != null) payload.amount = editForm.amount;
    await updateOrder(editTarget.value.id, payload as Partial<Order>);
    showEdit.value = false;
    editTarget.value = null;
    load();
  } catch (e: unknown) {
    errorMsg.value = e instanceof Error ? e.message : '更新失败';
  }
}

function confirmDelete(o: Order) {
  deleteTarget.value = o;
  showDelete.value = true;
}

async function doDelete() {
  if (!deleteTarget.value) return;
  try {
    await deleteOrder(deleteTarget.value.id);
    showDelete.value = false;
    deleteTarget.value = null;
    load();
  } catch (e: unknown) {
    errorMsg.value = e instanceof Error ? e.message : '删除失败';
  }
}

async function doExport() {
  try {
    const res = await exportOrders({
      materialId: query.materialId,
      specId: query.specId,
      vendorId: query.vendorId,
      orderNo: query.orderNo || undefined,
      dateFrom: query.dateFrom || undefined,
      dateTo: query.dateTo || undefined,
      sortBy: sortBy.value,
      sortOrder: sortOrder.value,
    });
    const filename = getExportFilename(res.headers['content-disposition'] as string);
    downloadBlob(new Blob([res.data]), filename);
  } catch (e: unknown) {
    errorMsg.value = e instanceof Error ? e.message : '导出失败';
  }
}

onMounted(() => {
  loadMeta();
  load();
});
</script>

<style scoped>
.order-list {
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

.input {
  padding: 0.5rem 0.75rem;
  border: 1px solid var(--color-border);
  border-radius: var(--radius-sm);
  min-width: 120px;
}

.actions {
  display: flex;
  gap: 0.5rem;
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

.data-table {
  width: 100%;
  border-collapse: collapse;
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

.form-row {
  display: flex;
  flex-wrap: wrap;
  gap: 1rem;
}

.form-row .field {
  min-width: 180px;
}

.form-row .field.full-width {
  min-width: 100%;
}

.required {
  color: var(--color-error);
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
