<template>
  <div class="space-y-8">
    <!-- Hero Section -->
    <div class="relative overflow-hidden rounded-[2rem] md:rounded-[2.5rem] bg-gradient-to-br from-rose-500 to-pink-600 text-white p-5 md:p-10 shadow-2xl mb-6 md:mb-10 sticky top-0 md:relative z-30">
      <!-- Background Decorative Elements -->
      <div class="absolute top-0 right-0 -mt-20 -mr-20 w-64 md:w-96 h-64 md:h-96 bg-white/10 rounded-full blur-3xl"></div>
      
      <div class="relative z-10">
        <div class="flex flex-col md:flex-row md:items-center justify-between gap-4 md:gap-6 mb-6 md:mb-8">
          <div>
            <h1 class="text-2xl md:text-4xl font-black tracking-tight mb-1">Habits</h1>
            <p class="text-rose-100 text-xs md:text-base font-medium opacity-90">Build consistency and track streaks.</p>
          </div>
          <button
            @click="showNewHabitModal = true"
            class="bg-white text-rose-600 px-4 md:px-6 py-2 md:py-3 rounded-xl md:rounded-2xl font-bold hover:bg-rose-50 transition-all flex items-center justify-center space-x-2 shadow-xl shadow-rose-900/20 group text-xs md:text-base"
          >
            <svg
              class="w-4 h-4 md:w-5 md:h-5 group-hover:rotate-90 transition-transform duration-300"
              fill="none"
              stroke="currentColor"
              viewBox="0 0 24 24"
            >
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2.5"
                d="M12 4v16m8-8H4"
              />
            </svg>
            <span>New Habit</span>
          </button>
        </div>

        <!-- Stats Bar -->
        <div class="flex items-center space-x-3 overflow-x-auto pb-2 md:pb-0 no-scrollbar">
          <div class="bg-white/10 backdrop-blur-md border border-white/20 px-3 md:px-5 py-1.5 md:py-2.5 rounded-xl md:rounded-2xl flex items-center space-x-2 md:space-x-3 whitespace-nowrap">
            <span class="text-base md:text-2xl font-black">{{ totalHabits }}</span>
            <span class="text-[8px] md:text-xs font-bold uppercase tracking-widest opacity-80">Total</span>
          </div>
          <div class="bg-white/10 backdrop-blur-md border border-white/20 px-3 md:px-5 py-1.5 md:py-2.5 rounded-xl md:rounded-2xl flex items-center space-x-2 md:space-x-3 whitespace-nowrap">
            <span class="text-base md:text-2xl font-black text-orange-300">{{ longestStreak }}</span>
            <span class="text-[8px] md:text-xs font-bold uppercase tracking-widest opacity-80">Streak</span>
          </div>
          <div class="bg-white/10 backdrop-blur-md border border-white/20 px-3 md:px-5 py-1.5 md:py-2.5 rounded-xl md:rounded-2xl flex items-center space-x-2 md:space-x-3 whitespace-nowrap">
            <span class="text-base md:text-2xl font-black text-emerald-300">{{ completedToday }}</span>
            <span class="text-[8px] md:text-xs font-bold uppercase tracking-widest opacity-80">Today</span>
          </div>
        </div>
      </div>
    </div>

    <!-- Habits Consistency Tracker -->
    <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
      <div v-if="habitStore.loading" class="col-span-full py-20 text-center">
        <div
          class="animate-spin rounded-full h-12 w-12 border-4 border-slate-100 border-t-indigo-600 mx-auto"
        ></div>
      </div>
      <div
        v-else
        v-for="habit in habitStore.habits"
        :key="habit.id"
        :id="`habit-${habit.id}`"
        class="bg-white dark:bg-slate-800 rounded-[2.5rem] border border-slate-100 dark:border-slate-700 shadow-sm hover:shadow-2xl hover:shadow-rose-100 dark:hover:shadow-none transition-all duration-500 p-6 md:p-6 group relative overflow-hidden flex flex-col"
      >
        <!-- Streak Background Decor -->
        <div class="absolute -right-4 -top-4 w-32 h-32 bg-orange-500/5 rounded-full blur-3xl group-hover:bg-orange-500/10 transition-colors"></div>

        <div class="flex items-start justify-between mb-6 relative z-10">
          <div class="flex items-center space-x-4">
            <div
              class="w-12 h-12 md:w-14 md:h-14 bg-rose-50 dark:bg-rose-900/20 text-rose-600 dark:text-rose-400 rounded-2xl flex items-center justify-center shadow-sm group-hover:scale-110 transition-transform duration-300"
            >
              <svg class="w-6 h-6 md:w-7 md:h-7" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4.318 6.318a4.5 4.5 0 000 6.364L12 20.364l7.682-7.682a4.5 4.5 0 00-6.364-6.364L12 7.636l-1.318-1.318a4.5 4.5 0 00-6.364 0z" />
              </svg>
            </div>
            <div>
              <h3 class="text-lg md:text-xl font-black text-slate-900 dark:text-white group-hover:text-rose-600 dark:group-hover:text-rose-400 transition-colors">
                {{ habit.name }}
              </h3>
              <p class="text-[10px] text-slate-500 dark:text-slate-400 font-bold uppercase tracking-[0.2em]">
                {{ habit.frequency }}
                <span v-if="habit.scheduled_days" class="ml-2 text-rose-500">• {{ formatScheduledDays(habit.scheduled_days) }}</span>
              </p>
            </div>
          </div>
          <div class="text-right">
            <div class="flex items-center justify-end space-x-1">
              <span class="text-3xl font-black text-slate-900 dark:text-white">{{ habit.current_streak || 0 }}</span>
              <span class="text-2xl">🔥</span>
            </div>
            <p class="text-[10px] font-black text-slate-400 dark:text-slate-500 uppercase tracking-widest">Day Streak</p>
          </div>
        </div>

        <!-- Daily Habit: Mini Heatmap Tracker -->
        <div v-if="habit.frequency === 'daily'" class="bg-slate-50 dark:bg-slate-900/50 rounded-2xl p-4 mb-6 relative z-10">
          <div class="flex items-center justify-between mb-3">
            <span class="text-[10px] font-black text-slate-400 dark:text-slate-500 uppercase tracking-widest">Weekly Consistency</span>
            <span class="text-[10px] font-bold text-indigo-600 dark:text-indigo-400">{{ calculateWeeklyCompletion(habit) }}%</span>
          </div>
          <div class="flex justify-between items-center px-1">
            <div
              v-for="day in last7Days"
              :key="day.date"
              class="flex flex-col items-center space-y-2"
            >
              <div
                @click="toggleHabit(habit.id, day.date)"
                class="w-8 h-8 rounded-lg border-2 transition-all cursor-pointer flex items-center justify-center"
                :class="[
                  isHabitCompleted(habit, day.date)
                    ? 'bg-emerald-500 border-emerald-500 text-white shadow-lg shadow-emerald-100'
                    : 'border-slate-200 dark:border-slate-700 bg-white dark:bg-slate-800 hover:border-emerald-300',
                ]"
              >
                <svg v-if="isHabitCompleted(habit, day.date)" class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="3" d="M5 13l4 4L19 7" />
                </svg>
              </div>
              <span class="text-[8px] font-black text-slate-400 dark:text-slate-500 uppercase">{{ day.label[0] }}</span>
            </div>
          </div>
        </div>

        <!-- Weekly Habit: Tasks List -->
        <div v-else class="mb-6 relative z-10">
          <button 
            @click="toggleExpand(habit.id)"
            class="w-full bg-slate-50 dark:bg-slate-900/50 rounded-2xl p-4 flex items-center justify-between group/btn hover:bg-slate-100 dark:hover:bg-slate-900 transition-all"
          >
            <div class="flex items-center gap-3">
              <div class="w-8 h-8 rounded-lg bg-rose-100 dark:bg-rose-900/30 flex items-center justify-center text-rose-600">
                <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5H7a2 2 0 00-2 2v12a2 2 0 002 2h10a2 2 0 002-2V7a2 2 0 00-2-2h-2M9 5a2 2 0 002 2h2a2 2 0 002-2M9 5a2 2 0 012-2h2a2 2 0 012 2" /></svg>
              </div>
              <span class="text-xs font-black text-slate-600 dark:text-slate-300 uppercase tracking-widest">Weekly Tasks</span>
            </div>
            <svg 
              class="w-4 h-4 text-slate-400 transition-transform duration-300" 
              :class="{ 'rotate-180': expandedHabitId === habit.id }"
              fill="none" stroke="currentColor" viewBox="0 0 24 24"
            >
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="M19 9l-7 7-7-7" />
            </svg>
          </button>

          <transition
            enter-active-class="transition duration-300 ease-out"
            enter-from-class="transform scale-95 opacity-0"
            enter-to-class="transform scale-100 opacity-100"
            leave-active-class="transition duration-200 ease-in"
            leave-from-class="transform scale-100 opacity-100"
            leave-to-class="transform scale-95 opacity-0"
          >
            <div v-if="expandedHabitId === habit.id" class="mt-3 space-y-2 px-2">
              <div 
                v-for="(task, idx) in getHabitTasks(habit.description)" 
                :key="idx"
                class="flex items-center gap-3 p-2 rounded-xl hover:bg-slate-50 dark:hover:bg-slate-800/50 transition-colors"
              >
                <div class="w-1.5 h-1.5 rounded-full bg-rose-400"></div>
                <span class="text-sm font-medium text-slate-600 dark:text-slate-400">{{ task }}</span>
              </div>
              <div v-if="getHabitTasks(habit.description).length === 0" class="py-4 text-center text-xs text-slate-400 italic">
                No tasks listed in description
              </div>
            </div>
          </transition>
        </div>

        <div v-if="isScheduledForToday(habit)" class="mt-auto flex items-center justify-between gap-4 relative z-10">
          <button
            @click="toggleHabit(habit.id, new Date().toISOString().split('T')[0])"
            class="flex-1 py-3 rounded-xl font-black text-sm transition-all flex items-center justify-center space-x-2"
            :class="[
              isHabitCompleted(habit, new Date().toISOString().split('T')[0])
                ? 'bg-slate-100 dark:bg-slate-700 text-slate-400 cursor-default'
                : 'bg-rose-600 text-white hover:bg-rose-700 shadow-lg shadow-rose-200 dark:shadow-none'
            ]"
          >
            <svg v-if="isHabitCompleted(habit, new Date().toISOString().split('T')[0])" class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="3" d="M5 13l4 4L19 7" />
            </svg>
            <span>{{ isHabitCompleted(habit, new Date().toISOString().split('T')[0]) ? 'Completed Today' : 'Check In Today' }}</span>
          </button>
          
          <div class="flex items-center space-x-1">
            <button
              @click.stop="openEditModal(habit)"
              class="p-2 text-slate-300 hover:text-indigo-500 transition-colors"
            >
              <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 5H6a2 2 0 00-2 2v12a2 2 0 002 2h10a2 2 0 002-2v-5M16.5 3.5a2.121 2.121 0 013 3L7 19l-4 1 1-4L16.5 3.5z" />
              </svg>
            </button>
            <button
              @click.stop="confirmDelete(habit)"
              class="p-2 text-slate-300 hover:text-rose-500 transition-colors"
            >
              <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16" />
              </svg>
            </button>
          </div>
        </div>
        <div v-else class="mt-auto flex items-center justify-between gap-4 relative z-10">
          <div class="flex-1 py-3 px-4 rounded-xl bg-slate-50 dark:bg-slate-900/50 border border-dashed border-slate-200 dark:border-slate-700 text-center">
            <span class="text-[10px] font-bold text-slate-400 uppercase tracking-widest">Not scheduled for today</span>
          </div>
          <div class="flex items-center space-x-1">
            <button
              @click.stop="openEditModal(habit)"
              class="p-2 text-slate-300 hover:text-indigo-500 transition-colors"
            >
              <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 5H6a2 2 0 00-2 2v12a2 2 0 002 2h10a2 2 0 002-2v-5M16.5 3.5a2.121 2.121 0 013 3L7 19l-4 1 1-4L16.5 3.5z" />
              </svg>
            </button>
            <button
              @click.stop="confirmDelete(habit)"
              class="p-2 text-slate-300 hover:text-rose-500 transition-colors"
            >
              <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16" />
              </svg>
            </button>
          </div>
        </div>
      </div>
    </div>

    <!-- Empty State -->
    <div
      v-if="!habitStore.loading && habitStore.habits.length === 0"
      class="py-20 text-center"
    >
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
            d="M4.318 6.318a4.5 4.5 0 000 6.364L12 20.364l7.682-7.682a4.5 4.5 0 00-6.364-6.364L12 7.636l-1.318-1.318a4.5 4.5 0 00-6.364 0z"
          />
        </svg>
      </div>
      <h3 class="text-xl font-bold text-slate-900 dark:text-white">
        No habits tracked
      </h3>
      <p class="text-slate-500 dark:text-slate-400 mt-2">
        Start building better routines today.
      </p>
    </div>

    <!-- New Habit Modal -->
    <div
      v-if="showNewHabitModal"
      class="fixed inset-0 bg-slate-900/50 backdrop-blur-sm z-50 flex items-center justify-center p-4"
    >
      <div
        class="bg-white dark:bg-slate-900 rounded-3xl w-full max-w-lg shadow-2xl overflow-hidden border border-slate-100 dark:border-slate-800"
      >
        <div
          class="p-6 border-b border-slate-50 dark:border-slate-800 flex justify-between items-center"
        >
          <h3 class="text-xl font-bold text-slate-900 dark:text-white">
            Create New Habit
          </h3>
          <button
            @click="showNewHabitModal = false"
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
        <div class="p-6 space-y-6 max-h-[60vh] overflow-y-auto custom-scrollbar">
          <div>
            <label class="block text-xs font-bold text-slate-400 uppercase tracking-widest mb-2">Habit Name</label>
            <input
              v-model="newHabit.name"
              type="text"
              class="w-full bg-slate-50 dark:bg-slate-800 border border-slate-100 dark:border-slate-700 rounded-xl px-4 py-3 outline-none focus:ring-2 focus:ring-indigo-500 text-slate-900 dark:text-white"
              placeholder="e.g. Morning Meditation"
            />
          </div>
          <div>
            <label class="block text-xs font-bold text-slate-400 uppercase tracking-widest mb-2">Frequency</label>
            <select
              v-model="newHabit.frequency"
              class="w-full bg-slate-50 dark:bg-slate-800 border border-slate-100 dark:border-slate-700 rounded-xl px-4 py-3 outline-none focus:ring-2 focus:ring-indigo-500 text-slate-900 dark:text-white"
            >
              <option value="daily">Daily</option>
              <option value="weekly">Weekly</option>
            </select>
          </div>
          <div v-if="newHabit.frequency === 'weekly'" class="space-y-6">
            <div>
              <label class="block text-xs font-bold text-slate-400 uppercase tracking-widest mb-3">Schedule Days</label>
              <div class="flex justify-between gap-1">
                <button
                  v-for="(day, index) in ['S', 'M', 'T', 'W', 'T', 'F', 'S']"
                  :key="index"
                  @click="toggleDayInNewHabit(index)"
                  class="flex-1 h-10 rounded-xl font-bold text-xs transition-all border-2"
                  :class="isDaySelectedInNewHabit(index) ? 'bg-rose-500 border-rose-500 text-white shadow-lg shadow-rose-100' : 'border-slate-100 dark:border-slate-800 text-slate-400 hover:border-rose-200'"
                >
                  {{ day }}
                </button>
              </div>
            </div>
            
            <div class="space-y-3">
              <label class="block text-xs font-bold text-slate-400 uppercase tracking-widest">Weekly Tasks</label>
              <div v-for="(task, idx) in newTaskItems" :key="idx" class="flex items-center gap-2">
                <input
                  v-model="newTaskItems[idx]"
                  type="text"
                  class="flex-1 bg-slate-50 dark:bg-slate-800 border border-slate-100 dark:border-slate-700 rounded-xl px-4 py-2 text-sm outline-none focus:ring-2 focus:ring-rose-500 text-slate-900 dark:text-white"
                  placeholder="Task name..."
                />
                <button @click="newTaskItems.splice(idx, 1)" class="p-2 text-slate-400 hover:text-rose-500">
                  <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" /></svg>
                </button>
              </div>
              <button 
                @click="newTaskItems.push('')"
                class="w-full py-2 border-2 border-dashed border-slate-200 dark:border-slate-700 rounded-xl text-xs font-bold text-slate-400 hover:border-rose-300 hover:text-rose-500 transition-all"
              >
                + Add Task
              </button>
            </div>
          </div>
        </div>
        <div
          class="p-6 bg-slate-50 dark:bg-slate-800/50 flex justify-end space-x-3"
        >
          <button
            @click="showNewHabitModal = false"
            class="px-6 py-2 text-slate-600 dark:text-slate-400 font-bold"
          >
            Cancel
          </button>
          <button
            @click="handleCreateHabit"
            class="px-6 py-2 bg-indigo-600 text-white rounded-xl font-bold hover:bg-indigo-700 transition-all shadow-lg shadow-indigo-100"
          >
            Create Habit
          </button>
        </div>
      </div>
    </div>

    <!-- Edit Habit Modal -->
    <div
      v-if="showEditModal"
      class="fixed inset-0 bg-slate-900/50 backdrop-blur-sm z-50 flex items-center justify-center p-4"
    >
      <div
        class="bg-white dark:bg-slate-900 rounded-3xl w-full max-w-lg shadow-2xl overflow-hidden border border-slate-100 dark:border-slate-800"
      >
        <div
          class="p-6 border-b border-slate-50 dark:border-slate-800 flex justify-between items-center"
        >
          <h3 class="text-xl font-bold text-slate-900 dark:text-white">
            Edit Habit
          </h3>
          <button
            @click="showEditModal = false"
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
        <div class="p-6 space-y-6 max-h-[60vh] overflow-y-auto custom-scrollbar">
          <div>
            <label class="block text-xs font-bold text-slate-400 uppercase tracking-widest mb-2">Habit Name</label>
            <input
              v-model="editHabitData.name"
              type="text"
              class="w-full bg-slate-50 dark:bg-slate-800 border border-slate-100 dark:border-slate-700 rounded-xl px-4 py-3 outline-none focus:ring-2 focus:ring-indigo-500 text-slate-900 dark:text-white"
            />
          </div>
          <div>
            <label class="block text-xs font-bold text-slate-400 uppercase tracking-widest mb-2">Frequency</label>
            <select
              v-model="editHabitData.frequency"
              class="w-full bg-slate-50 dark:bg-slate-800 border border-slate-100 dark:border-slate-700 rounded-xl px-4 py-3 outline-none focus:ring-2 focus:ring-indigo-500 text-slate-900 dark:text-white"
            >
              <option value="daily">Daily</option>
              <option value="weekly">Weekly</option>
            </select>
          </div>
          <div v-if="editHabitData.frequency === 'weekly'" class="space-y-6">
            <div>
              <label class="block text-xs font-bold text-slate-400 uppercase tracking-widest mb-3">Schedule Days</label>
              <div class="flex justify-between gap-1">
                <button
                  v-for="(day, index) in ['S', 'M', 'T', 'W', 'T', 'F', 'S']"
                  :key="index"
                  @click="toggleDayInEditHabit(index)"
                  class="flex-1 h-10 rounded-xl font-bold text-xs transition-all border-2"
                  :class="isDaySelectedInEditHabit(index) ? 'bg-rose-500 border-rose-500 text-white shadow-lg shadow-rose-100' : 'border-slate-100 dark:border-slate-800 text-slate-400 hover:border-rose-200'"
                >
                  {{ day }}
                </button>
              </div>
            </div>

            <div class="space-y-3">
              <label class="block text-xs font-bold text-slate-400 uppercase tracking-widest">Weekly Tasks</label>
              <div v-for="(task, idx) in editTaskItems" :key="idx" class="flex items-center gap-2">
                <input
                  v-model="editTaskItems[idx]"
                  type="text"
                  class="flex-1 bg-slate-50 dark:bg-slate-800 border border-slate-100 dark:border-slate-700 rounded-xl px-4 py-2 text-sm outline-none focus:ring-2 focus:ring-rose-500 text-slate-900 dark:text-white"
                  placeholder="Task name..."
                />
                <button @click="editTaskItems.splice(idx, 1)" class="p-2 text-slate-400 hover:text-rose-500">
                  <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" /></svg>
                </button>
              </div>
              <button 
                @click="editTaskItems.push('')"
                class="w-full py-2 border-2 border-dashed border-slate-200 dark:border-slate-700 rounded-xl text-xs font-bold text-slate-400 hover:border-rose-300 hover:text-rose-500 transition-all"
              >
                + Add Task
              </button>
            </div>
          </div>
        </div>
        <div
          class="p-6 bg-slate-50 dark:bg-slate-800/50 flex justify-end space-x-3"
        >
          <button
            @click="showEditModal = false"
            class="px-6 py-2 text-slate-600 dark:text-slate-400 font-bold"
          >
            Cancel
          </button>
          <button
            @click="handleUpdateHabit"
            class="px-6 py-2 bg-indigo-600 text-white rounded-xl font-bold hover:bg-indigo-700 transition-all shadow-lg shadow-indigo-100"
          >
            Save Changes
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted, inject, nextTick } from "vue";
import { useHabitStore } from "../stores/habits";
import { useRoute } from "vue-router";

const habitStore = useHabitStore();
const route = useRoute();
const showToast = inject("showToast");
const showNewHabitModal = ref(false);
const showEditModal = ref(false);
const expandedHabitId = ref(null);
const newHabit = ref({ name: "", frequency: "daily", scheduled_days: "" });
const editHabitData = ref({});
const newTaskItems = ref([]);
const editTaskItems = ref([]);

function formatScheduledDays(daysStr) {
  if (!daysStr) return "";
  const days = daysStr.split(",").map(Number);
  const dayNames = ["Sun", "Mon", "Tue", "Wed", "Thu", "Fri", "Sat"];
  return days.map(d => dayNames[d]).join(", ");
}

function toggleDayInNewHabit(dayIndex) {
  let days = newHabit.value.scheduled_days ? newHabit.value.scheduled_days.split(",") : [];
  const index = days.indexOf(dayIndex.toString());
  if (index === -1) {
    days.push(dayIndex.toString());
  } else {
    days.splice(index, 1);
  }
  newHabit.value.scheduled_days = days.sort().join(",");
}

function isDaySelectedInNewHabit(dayIndex) {
  if (!newHabit.value.scheduled_days) return false;
  return newHabit.value.scheduled_days.split(",").includes(dayIndex.toString());
}

function toggleDayInEditHabit(dayIndex) {
  let days = editHabitData.value.scheduled_days ? editHabitData.value.scheduled_days.split(",") : [];
  const index = days.indexOf(dayIndex.toString());
  if (index === -1) {
    days.push(dayIndex.toString());
  } else {
    days.splice(index, 1);
  }
  editHabitData.value.scheduled_days = days.sort().join(",");
}

function isDaySelectedInEditHabit(dayIndex) {
  if (!editHabitData.value.scheduled_days) return false;
  return editHabitData.value.scheduled_days.split(",").includes(dayIndex.toString());
}

function getHabitTasks(description) {
  if (!description) return [];
  return description
    .split("\n")
    .map(line => line.trim())
    .filter(line => line.startsWith("-") || line.startsWith("*"))
    .map(line => line.substring(1).trim());
}

function toggleExpand(id) {
  expandedHabitId.value = expandedHabitId.value === id ? null : id;
}

function isScheduledForToday(habit) {
  if (habit.frequency === 'daily') return true;
  if (!habit.scheduled_days) return false;
  const today = new Date().getDay().toString();
  return habit.scheduled_days.split(',').includes(today);
}

const totalHabits = computed(() => habitStore.habits.length);
const longestStreak = computed(() => {
  if (habitStore.habits.length === 0) return 0;
  return Math.max(...habitStore.habits.map(h => h.current_streak || 0));
});
const completedToday = computed(() => {
  const today = new Date().toISOString().split("T")[0];
  return habitStore.habits.filter(h => isHabitCompleted(h, today)).length;
});

const last7Days = computed(() => {
  const days = [];
  for (let i = 6; i >= 0; i--) {
    const d = new Date();
    d.setDate(d.getDate() - i);
    days.push({
      label: d.toLocaleDateString("en-US", { weekday: "short" }).charAt(0),
      date: d.toISOString().split("T")[0],
    });
  }
  return days;
});

function isHabitCompleted(habit, date) {
  if (!habit.logs) return false;
  return habit.logs.some((log) => log.log_date.startsWith(date));
}

function calculateWeeklyCompletion(habit) {
  const completedCount = last7Days.value.filter(day => isHabitCompleted(habit, day.date)).length;
  return Math.round((completedCount / 7) * 100);
}

async function toggleHabit(habitId, date) {
  try {
    await habitStore.logHabit(habitId, date);
    showToast("Habit logged!", "success");
  } catch (err) {
    showToast("Failed to log habit", "error");
    console.error(err);
  }
}

async function handleCreateHabit() {
  if (!newHabit.value.name) return;
  
  // Sync tasks to description
  if (newHabit.value.frequency === 'weekly' && newTaskItems.value.length > 0) {
    newHabit.value.description = newTaskItems.value
      .filter(t => t.trim())
      .map(t => `- ${t.trim()}`)
      .join('\n');
  }

  try {
    await habitStore.createHabit(newHabit.value);
    showNewHabitModal.value = false;
    newHabit.value = { name: "", frequency: "daily" };
    newTaskItems.value = [];
    showToast("Habit created!", "success");
  } catch (err) {
    showToast("Failed to create habit", "error");
    console.error(err);
  }
}

async function handleDeleteHabit(id) {
  if (confirm("Are you sure you want to delete this habit?")) {
    try {
      await habitStore.deleteHabit(id);
      showToast("Habit deleted", "success");
    } catch (err) {
      showToast("Failed to delete habit", "error");
    }
  }
}

function openEditModal(habit) {
  editHabitData.value = { ...habit };
  editTaskItems.value = getHabitTasks(habit.description);
  showEditModal.value = true;
}

async function handleUpdateHabit() {
  // Sync tasks to description
  if (editHabitData.value.frequency === 'weekly') {
    editHabitData.value.description = editTaskItems.value
      .filter(t => t.trim())
      .map(t => `- ${t.trim()}`)
      .join('\n');
  }

  try {
    await habitStore.updateHabit(editHabitData.value.id, editHabitData.value);
    showEditModal.value = false;
    showToast("Habit updated successfully", "success");
  } catch (err) {
    showToast("Failed to update habit", "error");
    console.error(err);
  }
}

onMounted(async () => {
  await habitStore.fetchHabits();
  
  // Handle navigation from Tasks page
  const habitId = route.query.id;
  const shouldExpand = route.query.expand === 'true';
  
  if (habitId) {
    nextTick(() => {
      const element = document.getElementById(`habit-${habitId}`);
      if (element) {
        element.scrollIntoView({ behavior: 'smooth', block: 'center' });
        if (shouldExpand) {
          expandedHabitId.value = parseInt(habitId);
        }
      }
    });
  }
});
</script>

<style scoped>
.custom-scrollbar::-webkit-scrollbar {
  width: 4px;
}
.custom-scrollbar::-webkit-scrollbar-track {
  background: transparent;
}
.custom-scrollbar::-webkit-scrollbar-thumb {
  background: #e2e8f0;
  border-radius: 10px;
}
.dark .custom-scrollbar::-webkit-scrollbar-thumb {
  background: #334155;
}
</style>
