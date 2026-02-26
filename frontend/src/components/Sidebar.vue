<template>
  <aside class="sidebar" :class="{ collapsed }">
    <div class="logo">
      <span class="logo-icon">📋</span>
      <span v-show="!collapsed" class="logo-text">钢管宝查询系统</span>
    </div>
    <nav class="nav">
      <div v-for="group in menuGroups" :key="group.title" class="nav-group">
        <div v-if="group.title && !collapsed" class="nav-group-title">
          {{ group.title }}
        </div>
        <router-link
          v-for="item in group.items"
          :key="item.name"
          :to="{ name: item.name }"
          class="nav-item"
          active-class="active"
        >
          <span class="nav-icon">{{ item.icon }}</span>
          <span v-show="!collapsed" class="nav-text">{{ item.title }}</span>
        </router-link>
      </div>
    </nav>
    <button
      class="collapse-btn"
      @click="collapsed = !collapsed"
      :title="collapsed ? '展开' : '收起'"
    >
      <span v-html="collapsed ? '▶' : '◀'"></span>
    </button>
  </aside>
</template>

<script setup lang="ts">
import { ref } from 'vue';

const collapsed = ref(false);

const menuGroups = [
  {
    title: '业务',
    items: [
      { name: 'Orders', title: '订单查询', icon: '🔍' },
      { name: 'Statistics', title: '统计分析', icon: '📊' },
    ],
  },
  {
    title: '基础数据',
    items: [
      { name: 'Materials', title: '材质管理', icon: '📦' },
      { name: 'Specs', title: '规格管理', icon: '📐' },
      { name: 'Vendors', title: '厂家管理', icon: '🏭' },
    ],
  },
];
</script>

<style scoped>
.sidebar {
  position: relative;
  width: 240px;
  min-height: 100vh;
  background: var(--color-sidebar);
  color: var(--color-sidebar-text);
  display: flex;
  flex-direction: column;
  transition: width 0.2s ease;
}

.sidebar.collapsed {
  width: 64px;
}

.logo {
  display: flex;
  align-items: center;
  gap: 0.75rem;
  padding: 1.25rem 1.25rem;
  border-bottom: 1px solid rgba(255, 255, 255, 0.08);
  flex-shrink: 0;
}

.logo-icon {
  font-size: 1.5rem;
  flex-shrink: 0;
}

.logo-text {
  font-size: 1rem;
  font-weight: 600;
  white-space: nowrap;
  overflow: hidden;
}

.nav {
  flex: 1;
  padding: 0.75rem 0;
  overflow-y: auto;
}

.nav-group {
  margin-bottom: 0.5rem;
}

.nav-group-title {
  padding: 0.5rem 1.25rem;
  font-size: 0.7rem;
  font-weight: 600;
  text-transform: uppercase;
  letter-spacing: 0.05em;
  color: var(--color-sidebar-group);
}

.nav-item {
  display: flex;
  align-items: center;
  gap: 0.75rem;
  padding: 0.6rem 1.25rem;
  color: var(--color-sidebar-text);
  opacity: 0.9;
  transition: all 0.15s;
}

.sidebar.collapsed .nav-item {
  justify-content: center;
  padding: 0.6rem;
}

.nav-item:hover {
  background: var(--color-sidebar-hover);
  color: var(--color-sidebar-text);
}

.nav-item.active {
  background: var(--color-sidebar-hover);
  color: var(--color-sidebar-text);
  border-left: 3px solid var(--color-sidebar-active);
  font-weight: 500;
}

.nav-icon {
  width: 20px;
  height: 20px;
  flex-shrink: 0;
  display: flex;
  align-items: center;
  justify-content: center;
}

.nav-icon {
  font-size: 1.1rem;
}

.nav-text {
  white-space: nowrap;
  overflow: hidden;
}

.collapse-btn {
  position: absolute;
  right: -12px;
  top: 50%;
  transform: translateY(-50%);
  width: 24px;
  height: 24px;
  border-radius: 50%;
  border: 1px solid var(--color-border);
  background: var(--color-canvas);
  color: var(--color-text);
  font-size: 10px;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  box-shadow: var(--shadow-sm);
  z-index: 10;
}

.collapse-btn:hover {
  background: var(--color-bg);
}
</style>
