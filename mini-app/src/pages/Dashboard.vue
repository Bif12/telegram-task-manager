<template>
  <div class="space-y-8">
    <!-- Hero Section -->
    <div class="relative overflow-hidden rounded-[2rem] md:rounded-[2.5rem] bg-gradient-to-br from-indigo-900 via-slate-900 to-slate-950 text-white p-5 md:p-10 shadow-2xl mb-6 md:mb-10 sticky top-0 md:relative z-30">
      <!-- Background Decorative Elements -->
      <div class="absolute top-0 right-0 -mt-20 -mr-20 w-64 md:w-96 h-64 md:h-96 bg-indigo-500/10 rounded-full blur-3xl"></div>
      <div class="absolute bottom-0 left-0 -mb-20 -ml-20 w-56 md:w-72 h-56 md:h-72 bg-blue-500/10 rounded-full blur-2xl"></div>
      
      <div class="relative z-10">
        <div class="flex flex-col md:flex-row md:items-center justify-between gap-4 md:gap-6 mb-6 md:mb-8">
          <div>
            <h1 class="text-xl md:text-4xl font-black tracking-tight mb-1">
              Welcome back, {{ user?.first_name || "User" }}! 👋
            </h1>
            <p class="text-slate-400 text-xs md:text-base font-medium opacity-90">Here's what's happening today.</p>
          </div>
          <div class="flex items-center space-x-2 md:space-x-3">
            <RouterLink
              to="/calendar"
              class="bg-white/10 backdrop-blur-md border border-white/20 text-white px-3 md:px-5 py-2 md:py-3 rounded-xl md:rounded-2xl font-bold hover:bg-white/20 transition-all flex items-center space-x-2 text-xs md:text-base"
            >
              <svg class="w-4 h-4 md:w-5 md:h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 7V3m8 4V3m-9 8h10M5 21h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v12a2 2 0 002 2z" />
              </svg>
              <span>Calendar</span>
            </RouterLink>
            <button
              @click="showNewTaskModal = true"
              class="bg-indigo-600 text-white px-3 md:px-6 py-2 md:py-3 rounded-xl md:rounded-2xl font-bold hover:bg-indigo-700 transition-all flex items-center space-x-2 shadow-xl shadow-indigo-900/40 group text-xs md:text-base"
            >
              <svg class="w-4 h-4 md:w-5 md:h-5 group-hover:rotate-90 transition-transform duration-300" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="M12 4v16m8-8H4" />
              </svg>
              <span>New Task</span>
            </button>
          </div>
        </div>

        <!-- Progress Overview (Cycle Show) -->
        <div class="flex flex-col md:flex-row items-center justify-center md:justify-start gap-6 md:gap-16">
          <!-- Main Progress Circle -->
          <div class="relative w-36 h-36 md:w-52 md:h-52 flex-shrink-0">
            <svg class="w-full h-full -rotate-90" viewBox="0 0 100 100">
              <!-- Background Circle -->
              <circle
                cx="50" cy="50" r="45"
                fill="none"
                stroke="currentColor"
                stroke-width="8"
                class="text-white/10"
              />
              <!-- Progress Circle -->
              <circle
                cx="50" cy="50" r="45"
                fill="none"
                stroke="currentColor"
                stroke-width="8"
                stroke-linecap="round"
                class="text-indigo-500 transition-all duration-1000 ease-out"
                :stroke-dasharray="282.7"
                :stroke-dashoffset="282.7 - (282.7 * overallCompletion) / 100"
              />
            </svg>
            <div class="absolute inset-0 flex flex-col items-center justify-center text-center">
              <span class="text-3xl md:text-5xl font-black">{{ Math.round(overallCompletion) }}%</span>
              <span class="text-[8px] md:text-xs font-bold uppercase tracking-widest opacity-60">Overall Done</span>
            </div>
          </div>

          <!-- Category Stats -->
          <div class="grid grid-cols-1 gap-3 w-full max-w-md">
            <div v-for="stat in dashboardStats" :key="stat.label" class="bg-white/5 backdrop-blur-md border border-white/10 p-3 md:p-4 rounded-2xl flex items-center space-x-4 group hover:bg-white/10 transition-all">
              <div :class="[stat.bg, stat.color]" class="w-8 h-8 md:w-10 md:h-10 rounded-xl flex items-center justify-center flex-shrink-0">
                <component :is="stat.icon" class="w-4 h-4 md:w-5 md:h-5" />
              </div>
              <div class="flex-1 min-w-0">
                <div class="flex items-center justify-between mb-1">
                  <span class="text-[8px] md:text-[10px] font-bold uppercase tracking-widest opacity-60 truncate">{{ stat.label }}</span>
                  <span class="text-[10px] md:text-xs font-bold" :class="stat.color">{{ stat.percent }}%</span>
                </div>
                <div class="w-full bg-white/10 h-1 md:h-1.5 rounded-full overflow-hidden">
                  <div 
                    class="h-full transition-all duration-1000" 
                    :class="stat.barColor"
                    :style="{ width: `${stat.percent}%` }"
                  ></div>
                </div>
                <div class="mt-1 text-sm md:text-lg font-black">{{ stat.value }}</div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <div class="grid grid-cols-1 lg:grid-cols-3 gap-8">
      <!-- Recent Tasks -->
      <div class="lg:col-span-2 space-y-6">
        <div class="flex items-center justify-between">
          <h2 class="text-xl font-bold text-slate-900 dark:text-white">
            Recent Tasks
          </h2>
          <RouterLink
            to="/tasks"
            class="text-indigo-600 text-sm font-semibold hover:text-indigo-700"
            >View all</RouterLink
          >
        </div>

        <div
          class="bg-white dark:bg-slate-900 rounded-2xl border border-slate-100 dark:border-slate-800 shadow-sm overflow-hidden"
        >
          <div v-if="taskStore.loading" class="p-12 text-center">
            <div
              class="animate-spin rounded-full h-10 w-10 border-4 border-slate-100 border-t-indigo-600 mx-auto"
            ></div>
          </div>
          <div
            v-else-if="taskStore.tasks.length === 0"
            class="p-12 text-center"
          >
            <div
              class="w-20 h-20 bg-slate-50 rounded-full flex items-center justify-center mx-auto mb-4"
            >
              <svg
                class="w-10 h-10 text-slate-300"
                fill="none"
                stroke="currentColor"
                viewBox="0 0 24 24"
              >
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                  d="M9 5H7a2 2 0 00-2 2v12a2 2 0 002 2h10a2 2 0 002-2V7a2 2 0 00-2-2h-2M9 5a2 2 0 002 2h2a2 2 0 002-2M9 5a2 2 0 012-2h2a2 2 0 012 2m-6 9l2 2 4-4"
                />
              </svg>
            </div>
            <h3 class="text-lg font-semibold text-slate-900 dark:text-white">
              No tasks yet
            </h3>
            <p class="text-slate-500 dark:text-slate-400 mt-1">
              Start by creating your first task above.
            </p>
          </div>
          <div v-else class="divide-y divide-slate-50 dark:divide-slate-800">
            <div
              v-for="task in recentTasks"
              :key="task.id"
              @click="$router.push(`/tasks/${task.slug}`)"
              class="p-4 hover:bg-slate-50 dark:hover:bg-slate-800 transition-all duration-200 flex items-center space-x-4 group cursor-pointer hover:pl-6"
            >
              <button
                @click="handleToggleTask(task)"
                class="w-6 h-6 rounded-full border-2 flex items-center justify-center transition-all"
                :class="[
                  task.status === 'completed'
                    ? 'bg-indigo-600 border-indigo-600'
                    : 'border-slate-200 dark:border-slate-700 hover:border-indigo-500',
                ]"
              >
                <svg
                  v-if="task.status === 'completed'"
                  class="w-4 h-4 text-white"
                  fill="none"
                  stroke="currentColor"
                  viewBox="0 0 24 24"
                >
                  <path
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    stroke-width="3"
                    d="M5 13l4 4L19 7"
                  />
                </svg>
                <div
                  v-else
                  class="w-3 h-3 rounded-full bg-indigo-500 opacity-0 group-hover:opacity-100 transition-opacity"
                ></div>
              </button>
              <div class="flex-1 min-w-0">
                <h4
                  :class="{
                    'line-through text-slate-400 dark:text-slate-500':
                      task.status === 'completed',
                  }"
                >
                  {{ task.title }}
                </h4>
                <p class="text-xs text-slate-500 dark:text-slate-400 truncate">
                  {{ task.project?.name || "No Project" }}
                </p>
              </div>
              <div class="flex items-center space-x-2">
                <span
                  :class="priorityClass(task.priority)"
                  class="text-[10px] font-bold px-2 py-0.5 rounded-full uppercase tracking-wider"
                >
                  {{ task.priority }}
                </span>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- Goals & Habits -->
      <div class="space-y-8">
        <!-- Projects Summary -->
        <section class="space-y-4">
          <div class="flex items-center justify-between">
            <h2 class="text-xl font-bold text-slate-900 dark:text-white">
              Projects
            </h2>
            <RouterLink
              to="/projects"
              class="text-indigo-600 text-sm font-semibold hover:text-indigo-700"
              >View all</RouterLink
            >
          </div>
          <div class="space-y-3">
            <div
              v-for="project in projectStore.projects.slice(0, 3)"
              :key="project.id"
              class="bg-white dark:bg-slate-900 p-4 rounded-2xl border border-slate-100 dark:border-slate-800 shadow-sm"
            >
              <h4 class="text-sm font-bold text-slate-800 dark:text-slate-200">
                {{ project.name }}
              </h4>
              <p class="text-xs text-slate-500 dark:text-slate-400 mt-1">
                {{ project.description }}
              </p>
            </div>
          </div>
        </section>
      </div>
    </div>

    <!-- New Task Modal (Simplified) -->
    <div
      v-if="showNewTaskModal"
      class="fixed inset-0 bg-slate-900/50 backdrop-blur-sm z-50 flex items-center justify-center p-4"
    >
      <div
        class="bg-white dark:bg-slate-900 rounded-3xl w-full max-w-lg shadow-2xl overflow-hidden border border-slate-100 dark:border-slate-800"
      >
        <div
          class="p-6 border-b border-slate-50 dark:border-slate-800 flex justify-between items-center"
        >
          <h3 class="text-xl font-bold text-slate-900 dark:text-white">
            Create New Task
          </h3>
          <button
            @click="showNewTaskModal = false"
            class="text-slate-400 hover:text-slate-600"
          >
            <svg
              class="w-6 h-6"
              fill="none"
              stroke="currentColor"
              viewBox="0 0 24 24"
            >
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
                d="M6 18L18 6M6 6l12 12"
              />
            </svg>
          </button>
        </div>
        <div class="p-6 space-y-4">
          <div>
            <label
              class="block text-xs font-bold text-slate-400 uppercase tracking-widest mb-2"
              >Task Title</label
            >
            <input
              v-model="newTask.title"
              type="text"
              class="w-full bg-slate-50 dark:bg-slate-800 border border-slate-100 dark:border-slate-700 rounded-xl px-4 py-3 outline-none focus:ring-2 focus:ring-indigo-500 text-slate-900 dark:text-white"
              placeholder="What needs to be done?"
            />
          </div>
          <div>
            <label
              class="block text-xs font-bold text-slate-400 uppercase tracking-widest mb-2"
              >Project</label
            >
            <select
              v-model="newTask.project_id"
              class="w-full bg-slate-50 dark:bg-slate-800 border border-slate-100 dark:border-slate-700 rounded-xl px-4 py-3 outline-none focus:ring-2 focus:ring-indigo-500 text-slate-900 dark:text-white"
            >
              <option :value="null">No Project</option>
              <option
                v-for="p in projectStore.projects"
                :key="p.id"
                :value="p.id"
              >
                {{ p.name }}
              </option>
            </select>
          </div>
          <div>
            <label
              class="block text-xs font-bold text-slate-400 uppercase tracking-widest mb-2"
              >Goal</label
            >
            <select
              v-model="newTask.goal_id"
              class="w-full bg-slate-50 dark:bg-slate-800 border border-slate-100 dark:border-slate-700 rounded-xl px-4 py-3 outline-none focus:ring-2 focus:ring-indigo-500 text-slate-900 dark:text-white"
            >
              <option :value="null">No Goal</option>
              <option v-for="g in goalStore.goals" :key="g.id" :value="g.id">
                {{ g.title }}
              </option>
            </select>
          </div>
          <div class="grid grid-cols-2 gap-4">
            <div>
              <label
                class="block text-xs font-bold text-slate-400 uppercase tracking-widest mb-2"
                >Priority</label
              >
              <select
                v-model="newTask.priority"
                class="w-full bg-slate-50 dark:bg-slate-800 border border-slate-100 dark:border-slate-700 rounded-xl px-4 py-3 outline-none focus:ring-2 focus:ring-indigo-500 text-slate-900 dark:text-white"
              >
                <option value="low">Low</option>
                <option value="medium">Medium</option>
                <option value="high">High</option>
              </select>
            </div>
          </div>
        </div>
        <div
          class="p-6 bg-slate-50 dark:bg-slate-800/50 flex justify-end space-x-3"
        >
          <button
            @click="showNewTaskModal = false"
            class="px-6 py-2 text-slate-600 dark:text-slate-400 font-bold"
          >
            Cancel
          </button>
          <button
            @click="handleCreateTask"
            class="px-6 py-2 bg-indigo-600 text-white rounded-xl font-bold hover:bg-indigo-700 transition-all shadow-lg shadow-indigo-100"
          >
            Create Task
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted, defineComponent, h, inject } from "vue";
import { useAuthStore } from "../stores/auth";
import { useTaskStore } from "../stores/tasks";
import { useProjectStore } from "../stores/projects";
import { useGoalStore } from "../stores/goals";

const authStore = useAuthStore();
const taskStore = useTaskStore();
const projectStore = useProjectStore();
const goalStore = useGoalStore();
const showToast = inject("showToast");

const user = computed(() => authStore.user);
const showNewTaskModal = ref(false);
const newTask = ref({
  title: "",
  project_id: null,
  goal_id: null,
  priority: "medium",
});

const recentTasks = computed(() => taskStore.tasks.slice(0, 5));

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
          d: "M9 5H7a2 2 0 00-2 2v12a2 2 0 002 2h10a2 2 0 002-2V7a2 2 0 00-2-2h-2M9 5a2 2 0 002 2h2a2 2 0 002-2M9 5a2 2 0 012-2h2a2 2 0 012 2m-6 9l2 2 4-4",
        }),
      ]
    ),
});

const IconClock = defineComponent({
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
          d: "M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z",
        }),
      ]
    ),
});

const dashboardStats = computed(() => {
  const totalTasks = taskStore.tasks.length;
  const completedTasks = taskStore.tasks.filter((t) => t.status === "completed").length;
  const taskPercent = totalTasks > 0 ? Math.round((completedTasks / totalTasks) * 100) : 0;

  const totalProjects = projectStore.projects.length;
  const completedProjects = projectStore.projects.filter(p => {
    const projectTasks = taskStore.tasks.filter(t => t.project_id === p.id);
    return projectTasks.length > 0 && projectTasks.every(t => t.status === 'completed');
  }).length;
  const projectPercent = totalProjects > 0 ? Math.round((completedProjects / totalProjects) * 100) : 0;

  const totalGoals = goalStore.goals.length;
  const completedGoals = goalStore.goals.filter(g => g.progress === 100).length;
  const goalPercent = totalGoals > 0 ? Math.round((completedGoals / totalGoals) * 100) : 0;

  return [
    {
      label: "Tasks",
      value: `${completedTasks}/${totalTasks}`,
      percent: taskPercent,
      icon: IconCheck,
      bg: "bg-emerald-500/10",
      color: "text-emerald-400",
      barColor: "bg-emerald-500",
    },
    {
      label: "Projects",
      value: `${completedProjects}/${totalProjects}`,
      percent: projectPercent,
      icon: IconClock,
      bg: "bg-indigo-500/10",
      color: "text-indigo-400",
      barColor: "bg-indigo-500",
    },
    {
      label: "Goals",
      value: `${completedGoals}/${totalGoals}`,
      percent: goalPercent,
      icon: IconClock,
      bg: "bg-amber-500/10",
      color: "text-amber-400",
      barColor: "bg-amber-500",
    }
  ];
});

const overallCompletion = computed(() => {
  const stats = dashboardStats.value;
  if (stats.length === 0) return 0;
  const totalPercent = stats.reduce((acc, curr) => acc + curr.percent, 0);
  return totalPercent / stats.length;
});

async function handleCreateTask() {
  if (!newTask.value.title) return;
  try {
    await taskStore.createTask(newTask.value);
    showNewTaskModal.value = false;
    newTask.value = {
      title: "",
      project_id: null,
      goal_id: null,
      priority: "medium",
    };
    showToast("Task created successfully", "success");
  } catch (err) {
    showToast("Failed to create task", "error");
    console.error(err);
  }
}

async function handleToggleTask(task) {
  try {
    await taskStore.toggleTaskStatus(task);
    showToast(
      task.status === "completed" ? "Task completed!" : "Task reopened",
      "success"
    );
  } catch (err) {
    showToast("Failed to update task", "error");
  }
}

const priorityClass = (priority) => {
  switch (priority) {
    case "high":
      return "bg-rose-50 text-rose-600";
    case "medium":
      return "bg-amber-50 text-amber-600";
    case "low":
      return "bg-emerald-50 text-emerald-600";
    default:
      return "bg-slate-50 text-slate-600";
  }
};

onMounted(() => {
  taskStore.fetchTasks();
  projectStore.fetchProjects();
  goalStore.fetchGoals();
});
</script>
