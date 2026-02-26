<template>
  <footer class="app-footer">
    <span class="copyright">{{ copyright }}</span>
  </footer>
</template>

<script setup lang="ts">
const copyright = '© 2025 规格及厂家查询系统 · 版权所有';
</script>

<style scoped>
.app-footer {
  flex-shrink: 0;
  padding: 0.75rem 1.5rem;
  text-align: center;
  font-size: 0.8rem;
  color: var(--color-text-muted);
  border-top: 1px solid var(--color-border);
  background: var(--color-bg);
}

.copyright {
  opacity: 0.85;
}
</style>
