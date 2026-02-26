<template>
  <Teleport to="body">
    <div v-if="modelValue" class="modal-overlay" @click.self="emit('update:modelValue', false)">
      <div class="modal">
        <div class="modal-header">
          <h3>{{ title }}</h3>
          <button class="modal-close" type="button" @click="emit('update:modelValue', false)">
            ×
          </button>
        </div>
        <div class="modal-body">
          <slot />
        </div>
        <div v-if="$slots.footer" class="modal-footer">
          <slot name="footer" />
        </div>
      </div>
    </div>
  </Teleport>
</template>

<script setup lang="ts">
defineProps<{ title: string; modelValue: boolean }>();
const emit = defineEmits<{ 'update:modelValue': [v: boolean] }>();
</script>

<style scoped>
.modal-overlay {
  position: fixed;
  inset: 0;
  background: rgba(0, 0, 0, 0.4);
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 1000;
}

.modal {
  background: var(--color-canvas);
  border: 1px solid var(--color-border);
  border-radius: var(--radius);
  box-shadow: 0 8px 24px rgba(31, 35, 40, 0.12);
  min-width: 360px;
  max-width: 90vw;
  max-height: 90vh;
  display: flex;
  flex-direction: column;
}

.modal-header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 1rem 1.25rem;
  border-bottom: 1px solid var(--color-border);
}

.modal-header h3 {
  font-size: 1.1rem;
  font-weight: 600;
}

.modal-close {
  width: 28px;
  height: 28px;
  border: none;
  background: transparent;
  font-size: 1.5rem;
  line-height: 1;
  color: var(--color-text-muted);
  cursor: pointer;
}

.modal-close:hover {
  color: var(--color-text);
}

.modal-body {
  padding: 1.25rem;
  overflow-y: auto;
}

.modal-footer {
  padding: 1rem 1.25rem;
  border-top: 1px solid var(--color-border);
  display: flex;
  justify-content: flex-end;
  gap: 0.5rem;
}
</style>
