<template>
  <div class="breadcrumb">
    <template v-for="(item, i) in items" :key="i">
      <router-link
        v-if="i < items.length - 1"
        :to="i === 0 ? { name: 'Orders' } : undefined"
        class="crumb"
      >
        {{ item }}
      </router-link>
      <span v-else class="crumb last">{{ item }}</span>
      <span v-if="i < items.length - 1" class="separator">/</span>
    </template>
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue';
import { useRoute } from 'vue-router';

const route = useRoute();

const items = computed(() => {
  const meta = route.meta?.breadcrumb;
  if (Array.isArray(meta)) return meta;
  return [route.meta?.title ?? '首页'];
});
</script>

<style scoped>
.breadcrumb {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  padding: 0.75rem 1.5rem;
  background: var(--color-canvas);
  border-bottom: 1px solid var(--color-border);
  font-size: 0.875rem;
}

.crumb {
  color: var(--color-text-muted);
}

.crumb:not(.last):hover {
  color: var(--color-accent);
}

.separator {
  margin: 0 0.5rem;
  color: var(--color-text-muted);
}

.crumb.last {
  color: var(--color-text);
  font-weight: 500;
  cursor: default;
}
</style>
