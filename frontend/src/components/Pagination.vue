<template>
  <div class="pagination">
    <div class="pagination-info">
      共 {{ total }} 条，
      <select v-model="selectValue" @change="onSelectChange" class="page-size-select">
        <option v-for="n in presetSizes" :key="n" :value="n">{{ n }} 条/页</option>
        <option v-if="allowCustom" value="custom">自定义</option>
      </select>
      <template v-if="allowCustom && selectValue === 'custom'">
        <input
          v-model.number="customInput"
          type="number"
          min="1"
          max="500"
          class="custom-input"
          placeholder="输入条数"
          @blur="applyCustom"
          @keyup.enter="applyCustom"
        />
      </template>
    </div>
    <div class="pagination-btns">
      <button class="btn" :disabled="page <= 1" @click="goTo(1)">首页</button>
      <button class="btn" :disabled="page <= 1" @click="goTo(page - 1)">上一页</button>
      <span class="page-num"> 第 {{ page }} / {{ totalPages }} 页 </span>
      <button class="btn" :disabled="page >= totalPages" @click="goTo(page + 1)">下一页</button>
      <button class="btn" :disabled="page >= totalPages" @click="goTo(totalPages)">末页</button>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, watch } from 'vue';

const props = withDefaults(
  defineProps<{
    total: number;
    page: number;
    pageSize: number;
    presetSizes?: number[];
    allowCustom?: boolean;
  }>(),
  {
    presetSizes: () => [10, 20, 30, 40, 50, 100],
    allowCustom: true,
  }
);

const emit = defineEmits<{
  'update:page': [value: number];
  'update:pageSize': [value: number];
}>();

const selectValue = ref<number | 'custom'>(props.pageSize);
const customInput = ref('');

const effectivePageSize = computed(() => {
  if (selectValue.value === 'custom') {
    const n = Number(customInput.value);
    return n >= 1 && n <= 500 ? n : props.pageSize;
  }
  return (selectValue.value as number) || props.pageSize;
});

const totalPages = computed(() =>
  props.total > 0 ? Math.ceil(props.total / effectivePageSize.value) : 1
);

watch(
  () => props.pageSize,
  (v) => {
    if (props.presetSizes.includes(v)) {
      selectValue.value = v;
    } else {
      selectValue.value = 'custom';
      customInput.value = String(v);
    }
  },
  { immediate: true }
);

function onSelectChange() {
  if (selectValue.value !== 'custom') {
    emit('update:pageSize', selectValue.value as number);
    emit('update:page', 1);
  }
}

function applyCustom() {
  const n = Number(customInput.value);
  if (n >= 1 && n <= 500) {
    emit('update:pageSize', n);
    emit('update:page', 1);
  } else {
    customInput.value = String(props.pageSize);
  }
}

function goTo(p: number) {
  if (p >= 1 && p <= totalPages.value) {
    emit('update:page', p);
  }
}
</script>

<style scoped>
.pagination {
  display: flex;
  align-items: center;
  justify-content: space-between;
  flex-wrap: wrap;
  gap: 1rem;
  padding: 0.75rem 0;
}

.pagination-info {
  display: flex;
  align-items: center;
  gap: 0.5rem;
}

.page-size-select {
  padding: 0.35rem 0.5rem;
  border: 1px solid var(--color-border);
  border-radius: 4px;
  background: var(--color-canvas);
}

.custom-input {
  width: 70px;
  padding: 0.35rem 0.5rem;
  border: 1px solid var(--color-border);
  border-radius: 4px;
}

.pagination-btns {
  display: flex;
  align-items: center;
  gap: 0.5rem;
}

.btn {
  padding: 0.4rem 0.75rem;
  border: 1px solid var(--color-border);
  border-radius: 4px;
  background: var(--color-canvas);
  font-size: 0.9rem;
}

.btn:hover:not(:disabled) {
  background: var(--color-bg);
  border-color: var(--color-accent);
}

.btn:disabled {
  opacity: 0.5;
  cursor: not-allowed;
}

.page-num {
  font-size: 0.9rem;
  color: var(--color-text-muted);
}
</style>
