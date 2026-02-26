<template>
  <div class="app-layout">
    <Sidebar class="sidebar" />
    <main class="main">
      <Breadcrumb class="breadcrumb" />
      <div class="content-area">
        <router-view v-slot="{ Component }">
          <transition name="fade" mode="out-in">
            <component :is="Component" />
          </transition>
        </router-view>
      </div>
      <Footer class="app-footer" />
    </main>
  </div>
</template>

<script setup lang="ts">
import Sidebar from './components/Sidebar.vue';
import Breadcrumb from './components/Breadcrumb.vue';
import Footer from './components/Footer.vue';
</script>

<style scoped>
.app-layout {
  display: flex;
  min-height: 100vh;
}

.sidebar {
  flex-shrink: 0;
}

.main {
  flex: 1;
  display: flex;
  flex-direction: column;
  overflow: hidden;
}

.breadcrumb {
  flex-shrink: 0;
}

.content-area {
  flex: 1;
  overflow: auto;
  padding: 1.25rem 1.5rem;
  min-height: 0;
}

.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.15s ease;
}
.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}
</style>
