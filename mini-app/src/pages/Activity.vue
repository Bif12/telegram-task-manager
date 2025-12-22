<template>
  <div class="space-y-8">
    <!-- Header -->
    <div
      class="flex flex-col md:flex-row md:items-center justify-between gap-4"
    >
      <div>
        <h1 class="text-3xl font-bold text-slate-900 dark:text-white">
          Activity Log
        </h1>
        <p class="text-slate-500 dark:text-slate-400 mt-1">
          Keep track of all changes and updates.
        </p>
      </div>
    </div>

    <!-- Activity Timeline -->
    <div
      class="space-y-8 relative before:absolute before:inset-0 before:left-8 before:w-0.5 before:bg-slate-100 dark:before:bg-slate-800 before:hidden md:before:block"
    >
      <div v-if="loading" class="py-20 text-center">
        <div
          class="animate-spin rounded-full h-12 w-12 border-4 border-slate-100 border-t-indigo-600 mx-auto"
        ></div>
      </div>

      <div v-else-if="activities.length === 0" class="py-20 text-center">
        <div
          class="w-24 h-24 bg-slate-100 dark:bg-slate-800 rounded-full flex items-center justify-center mx-auto mb-6"
        >
          <svg
            class="w-12 h-12 text-slate-300"
            fill="none"
            stroke="currentColor"
            viewBox="0 0 24 24"
          >
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z"
            />
          </svg>
        </div>
        <h3 class="text-xl font-bold text-slate-900 dark:text-white">
          No activity yet
        </h3>
        <p class="text-slate-500 dark:text-slate-400 mt-2">
          Your actions will appear here as you use the app.
        </p>
      </div>

      <div
        v-else
        v-for="(group, date) in groupedActivities"
        :key="date"
        class="space-y-4"
      >
        <div class="md:sticky md:top-24 z-10">
          <span
            class="bg-slate-900 dark:bg-indigo-600 text-white text-[10px] font-black px-3 py-1.5 rounded-full uppercase tracking-[0.2em] shadow-lg"
          >
            {{ date }}
          </span>
        </div>

        <div class="grid grid-cols-1 gap-4">
          <div
            v-for="activity in group"
            :key="activity.id"
            class="bg-white dark:bg-slate-900 p-5 rounded-2xl border border-slate-100 dark:border-slate-800 shadow-sm hover:shadow-md transition-all flex items-center space-x-4 relative md:ml-16"
          >
            <!-- Timeline Dot (Desktop) -->
            <div
              class="hidden md:block absolute -left-[41px] w-4 h-4 rounded-full border-4 border-white dark:border-slate-900 bg-indigo-600 shadow-sm"
            ></div>

            <div
              :class="actionColor(activity.action)"
              class="w-12 h-12 rounded-xl flex items-center justify-center flex-shrink-0"
            >
              <component :is="actionIcon(activity.action)" class="w-6 h-6" />
            </div>

            <div class="flex-1 min-w-0">
              <div class="flex items-center justify-between mb-1">
                <h3
                  class="text-sm font-bold text-slate-800 dark:text-slate-200"
                >
                  {{ activity.action }} {{ activity.resource_type }}
                </h3>
                <span
                  class="text-[10px] font-bold text-slate-400 dark:text-slate-500 uppercase tracking-widest"
                >
                  {{ formatTime(activity.created_at) }}
                </span>
              </div>
              <p class="text-xs text-slate-500 dark:text-slate-400 truncate">
                {{ activity.description || "No additional details" }}
              </p>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted, defineComponent, h } from "vue";
import api from "../services/api";

const activities = ref([]);
const loading = ref(true);

const groupedActivities = computed(() => {
  const groups = {};
  activities.value.forEach((activity) => {
    const date = new Date(activity.created_at).toLocaleDateString("en-US", {
      month: "long",
      day: "numeric",
      year: "numeric",
    });
    if (!groups[date]) groups[date] = [];
    groups[date].push(activity);
  });
  return groups;
});

async function fetchActivities() {
  loading.value = true;
  try {
    const response = await api.get("/activities");
    activities.value = response.data;
  } catch (err) {
    console.error(err);
  } finally {
    loading.value = false;
  }
}

const IconPlus = defineComponent({
  render: () =>
    h(
      "svg",
      {
        fill: "none",
        stroke: "currentColor",
        viewBox: "0 0 24 24",
        class: "w-inherit h-inherit",
      },
      [
        h("path", {
          "stroke-linecap": "round",
          "stroke-linejoin": "round",
          "stroke-width": "2",
          d: "M12 4v16m8-8H4",
        }),
      ]
    ),
});

const IconCheck = defineComponent({
  render: () =>
    h(
      "svg",
      {
        fill: "none",
        stroke: "currentColor",
        viewBox: "0 0 24 24",
        class: "w-inherit h-inherit",
      },
      [
        h("path", {
          "stroke-linecap": "round",
          "stroke-linejoin": "round",
          "stroke-width": "2",
          d: "M5 13l4 4L19 7",
        }),
      ]
    ),
});

const IconEdit = defineComponent({
  render: () =>
    h(
      "svg",
      {
        fill: "none",
        stroke: "currentColor",
        viewBox: "0 0 24 24",
        class: "w-inherit h-inherit",
      },
      [
        h("path", {
          "stroke-linecap": "round",
          "stroke-linejoin": "round",
          "stroke-width": "2",
          d: "M11 5H6a2 2 0 00-2 2v12a2 2 0 002 2h10a2 2 0 002-2v-5M16.5 3.5a2.121 2.121 0 013 3L7 19l-4 1 1-4L16.5 3.5z",
        }),
      ]
    ),
});

function actionIcon(action) {
  if (action.includes("Create")) return IconPlus;
  if (action.includes("Complete")) return IconCheck;
  return IconEdit;
}

function actionColor(action) {
  if (action.includes("Create"))
    return "bg-emerald-50 dark:bg-emerald-900/20 text-emerald-600 dark:text-emerald-400";
  if (action.includes("Complete"))
    return "bg-indigo-50 dark:bg-indigo-900/20 text-indigo-600 dark:text-indigo-400";
  return "bg-amber-50 dark:bg-amber-900/20 text-amber-600 dark:text-amber-400";
}

function formatTime(dateStr) {
  return new Date(dateStr).toLocaleTimeString("en-US", {
    hour: "2-digit",
    minute: "2-digit",
  });
}

onMounted(fetchActivities);
</script>
