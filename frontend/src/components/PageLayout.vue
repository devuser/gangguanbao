<template>
  <div class="page-layout">
    <header class="page-header">
      <div class="header-main">
        <h1 class="page-title">{{ title }}</h1>
        <p v-if="desc" class="page-desc">{{ desc }}</p>
      </div>
      <div v-if="$slots.actions" class="header-actions">
        <slot name="actions" />
      </div>
    </header>
    <div class="page-content">
      <slot />
    </div>
  </div>
</template>

<script setup lang="ts">
defineProps<{
  title: string;
  desc?: string;
}>();
</script>

<style scoped>
.page-layout {
  display: flex;
  flex-direction: column;
  height: 100%;
  min-height: 0;
}

.page-header {
  display: flex;
  align-items: flex-start;
  justify-content: space-between;
  gap: 1rem;
  margin-bottom: 1.25rem;
  flex-shrink: 0;
}

.header-main {
  flex: 1;
  min-width: 0;
}

.page-title {
  font-size: 1.35rem;
  font-weight: 600;
  color: var(--color-text);
  margin-bottom: 0.25rem;
}

.page-desc {
  font-size: 0.9rem;
  color: var(--color-text-muted);
}

.header-actions {
  flex-shrink: 0;
}

.page-content {
  flex: 1;
  min-height: 0;
  display: flex;
  flex-direction: column;
}
</style>
