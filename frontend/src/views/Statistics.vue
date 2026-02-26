<template>
  <PageLayout title="统计分析" desc="数据概览与订单趋势分析">
    <div class="statistics">
      <div v-if="errorMsg" class="error-msg">{{ errorMsg }}</div>

      <section class="stats-cards">
        <div class="card stat-card">
          <span class="stat-label">订单条数</span>
          <span class="stat-value">{{ stats?.counts?.orderCount ?? '-' }}</span>
        </div>
        <div class="card stat-card">
          <span class="stat-label">材质条数</span>
          <span class="stat-value">{{ stats?.counts?.materialCount ?? '-' }}</span>
        </div>
        <div class="card stat-card">
          <span class="stat-label">规格条数</span>
          <span class="stat-value">{{ stats?.counts?.specCount ?? '-' }}</span>
        </div>
        <div class="card stat-card">
          <span class="stat-label">厂家条数</span>
          <span class="stat-value">{{ stats?.counts?.vendorCount ?? '-' }}</span>
        </div>
      </section>

      <section class="chart-section">
        <div class="card chart-card">
          <h3 class="chart-title">每年订单数</h3>
          <div ref="chartYearRef" class="chart"></div>
        </div>
      </section>

      <section class="chart-section chart-row">
        <div class="card chart-card">
          <h3 class="chart-title">{{ curYear }}年各月订单数</h3>
          <div ref="chartMonthCurRef" class="chart"></div>
        </div>
        <div class="card chart-card">
          <h3 class="chart-title">{{ lastYear }}年各月订单数</h3>
          <div ref="chartMonthLastRef" class="chart"></div>
        </div>
      </section>

      <section class="chart-section chart-row">
        <div class="card chart-card">
          <h3 class="chart-title">{{ curYear }}年各材质订单数</h3>
          <div ref="chartMaterialRef" class="chart"></div>
        </div>
        <div class="card chart-card">
          <h3 class="chart-title">{{ curYear }}年各规格订单数</h3>
          <div ref="chartSpecRef" class="chart"></div>
        </div>
      </section>
    </div>
  </PageLayout>
</template>

<script setup lang="ts">
import { ref, onMounted, onUnmounted, nextTick } from 'vue';
import * as echarts from 'echarts';
import type { ECharts } from 'echarts';
import { fetchStats, type StatsResponse } from '@/api/stats';
import PageLayout from '@/components/PageLayout.vue';

const stats = ref<StatsResponse | null>(null);
const errorMsg = ref('');
const curYear = new Date().getFullYear();
const lastYear = curYear - 1;

const chartYearRef = ref<HTMLElement | null>(null);
const chartMonthCurRef = ref<HTMLElement | null>(null);
const chartMonthLastRef = ref<HTMLElement | null>(null);
const chartMaterialRef = ref<HTMLElement | null>(null);
const chartSpecRef = ref<HTMLElement | null>(null);

let chartYear: ECharts | null = null;
let chartMonthCur: ECharts | null = null;
let chartMonthLast: ECharts | null = null;
let chartMaterial: ECharts | null = null;
let chartSpec: ECharts | null = null;

const monthLabels = [
  '1月',
  '2月',
  '3月',
  '4月',
  '5月',
  '6月',
  '7月',
  '8月',
  '9月',
  '10月',
  '11月',
  '12月',
];

function initCharts() {
  if (!stats.value) return;
  disposeCharts();

  if (chartYearRef.value) {
    chartYear = echarts.init(chartYearRef.value);
    const yearData = stats.value.ordersByYear ?? [];
    chartYear.setOption({
      tooltip: { trigger: 'axis' },
      xAxis: {
        type: 'category',
        data: yearData.length ? yearData.map((x) => String(x.year)) : ['暂无数据'],
      },
      yAxis: { type: 'value', name: '订单数' },
      series: [
        {
          type: 'bar',
          data: yearData.length ? yearData.map((x) => x.count) : [0],
          itemStyle: { color: '#0969da' },
        },
      ],
    });
  }

  if (chartMonthCurRef.value) {
    chartMonthCur = echarts.init(chartMonthCurRef.value);
    chartMonthCur.setOption({
      tooltip: { trigger: 'axis' },
      xAxis: { type: 'category', data: monthLabels },
      yAxis: { type: 'value', name: '订单数' },
      series: [
        {
          type: 'bar',
          data: stats.value.ordersByMonthCur?.map((x) => x.count) ?? Array(12).fill(0),
          itemStyle: { color: '#1a7f37' },
        },
      ],
    });
  }

  if (chartMonthLastRef.value) {
    chartMonthLast = echarts.init(chartMonthLastRef.value);
    chartMonthLast.setOption({
      tooltip: { trigger: 'axis' },
      xAxis: { type: 'category', data: monthLabels },
      yAxis: { type: 'value', name: '订单数' },
      series: [
        {
          type: 'bar',
          data: stats.value.ordersByMonthLast?.map((x) => x.count) ?? Array(12).fill(0),
          itemStyle: { color: '#57606a' },
        },
      ],
    });
  }

  if (chartMaterialRef.value) {
    chartMaterial = echarts.init(chartMaterialRef.value);
    const materialData = stats.value.ordersByMaterial ?? [];
    chartMaterial.setOption({
      tooltip: { trigger: 'item' },
      series: [
        {
          type: 'pie',
          radius: ['40%', '70%'],
          data: materialData.length
            ? materialData.map((x) => ({ name: x.name, value: x.count }))
            : [{ name: '暂无数据', value: 1 }],
          emphasis: {
            itemStyle: { shadowBlur: 10, shadowOffsetX: 0, shadowColor: 'rgba(0,0,0,0.2)' },
          },
        },
      ],
    });
  }

  if (chartSpecRef.value) {
    chartSpec = echarts.init(chartSpecRef.value);
    const specData = stats.value.ordersBySpec ?? [];
    chartSpec.setOption({
      tooltip: { trigger: 'item' },
      series: [
        {
          type: 'pie',
          radius: ['40%', '70%'],
          data: specData.length
            ? specData.map((x) => ({ name: x.name, value: x.count }))
            : [{ name: '暂无数据', value: 1 }],
          emphasis: {
            itemStyle: { shadowBlur: 10, shadowOffsetX: 0, shadowColor: 'rgba(0,0,0,0.2)' },
          },
        },
      ],
    });
  }
}

function resizeCharts() {
  chartYear?.resize();
  chartMonthCur?.resize();
  chartMonthLast?.resize();
  chartMaterial?.resize();
  chartSpec?.resize();
}

function disposeCharts() {
  chartYear?.dispose();
  chartMonthCur?.dispose();
  chartMonthLast?.dispose();
  chartMaterial?.dispose();
  chartSpec?.dispose();
  chartYear = chartMonthCur = chartMonthLast = chartMaterial = chartSpec = null;
}

async function load() {
  errorMsg.value = '';
  try {
    const res = await fetchStats();
    stats.value = res.data;
    await nextTick();
    initCharts();
  } catch (e: unknown) {
    errorMsg.value = e instanceof Error ? e.message : '加载失败';
  }
}

onMounted(() => {
  load();
  window.addEventListener('resize', resizeCharts);
});

onUnmounted(() => {
  window.removeEventListener('resize', resizeCharts);
  disposeCharts();
});
</script>

<style scoped>
.statistics {
  display: flex;
  flex-direction: column;
  gap: 1.5rem;
  flex: 1;
  min-height: 0;
}

.error-msg {
  padding: 0.75rem 1rem;
  background: #ffebe9;
  color: var(--color-error);
  border-radius: var(--radius-sm);
  font-size: 0.9rem;
}

.stats-cards {
  display: grid;
  grid-template-columns: repeat(4, 1fr);
  gap: 1rem;
}

.stat-card {
  padding: 1.25rem;
  text-align: center;
}

.stat-label {
  display: block;
  font-size: 0.85rem;
  color: var(--color-text-muted);
  margin-bottom: 0.5rem;
}

.stat-value {
  font-size: 1.75rem;
  font-weight: 600;
  color: var(--color-text);
}

.chart-section {
  display: flex;
  flex-direction: column;
}

.chart-row {
  flex-direction: row;
  gap: 1rem;
}

.chart-row .chart-card {
  flex: 1;
  min-width: 0;
}

.chart-card {
  padding: 1.25rem;
}

.chart-title {
  font-size: 1rem;
  font-weight: 600;
  margin-bottom: 1rem;
  color: var(--color-text);
}

.chart {
  width: 100%;
  height: 280px;
}

@media (max-width: 900px) {
  .stats-cards {
    grid-template-columns: repeat(2, 1fr);
  }

  .chart-row {
    flex-direction: column;
  }
}
</style>
