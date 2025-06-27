<script setup lang="ts">
import { ref, onMounted } from "vue";

import QuakeMap from "@/components/QuakeMap.vue";

const eventId = ref("");
const isDebug = ref(false);
const isLoading = ref(true);
const currentEventIndex = ref(0);
const isSeekableToNext = ref(false);
const isSeekableToPrev = ref(false);

let events: string[] | undefined = undefined;

const fetchEventId = async (index: number): Promise<{id: string, isLatest: boolean, isFinal: boolean}> => {
    if (!events) {
        throw new Error("Events not loaded");
    }

    // IDに_VXSE51または_VXSE53が含まれるもののうちindex番目のイベントを取得
    const filtered = events.filter((id) => id.includes("_VXSE51") || id.includes("_VXSE53"));
    if (index < 0 || index >= filtered.length) {
        throw new Error("Index out of bounds");
    }

    return { id: filtered[index], isLatest: index === 0, isFinal: index === filtered.length - 1 };
};

const updateEventId = async (index: number): Promise<void> => {
    console.log(`Updating event ID to index: ${index}`);
    const fetched = await fetchEventId(index);
    eventId.value = fetched.id;
    isSeekableToNext.value = !fetched.isLatest;
    isSeekableToPrev.value = !fetched.isFinal;

    // URLパラメーターを追加
    const url = new URL(window.location.href);
    if (index === 0) {
        url.searchParams.delete("eventId");
    } else {
        url.searchParams.set("eventId", fetched.id);
    }

    window.history.replaceState({}, "", url.toString());
};

const shareCurrentURL = async (): Promise<void> => {
    const url = new URL(window.location.href);

    // Firefoxでは　navigator.shareがサポートされていないため、必要に応じてClipboard APIにフォールバック
    "share" in navigator ? navigator.share({ url: url.toString() }) : (navigator as Navigator).clipboard.writeText(url.toString());
};

onMounted(async () => {
    if (window.localStorage.getItem("DBG_DUMMY_EVENT_ID")) {
        isDebug.value = true;
        eventId.value = window.localStorage.getItem("DBG_DUMMY_EVENT_ID")!;
        isLoading.value = false;

        return;
    }

    const resp = await fetch("https://quake-jade.vercel.app/api/events");
    events = (await resp.json() as { events: string[] }).events;

    // eventIdがURLパラメーターに含まれている場合はそれを優先
    const urlParams = new URLSearchParams(window.location.search);
    const eventIdFromUrl = urlParams.get("eventId");
    if (eventIdFromUrl) {
        const index = events.indexOf(eventIdFromUrl);
        if (index !== -1) {
            currentEventIndex.value = index;
        } else {
            console.warn(`Event ID ${eventIdFromUrl} not found in events list.`);
        }
    }

    await updateEventId(currentEventIndex.value);

    isLoading.value = false;
});
</script>

<template>
    <main>
        <QuakeMap v-if=!isLoading :is-debug="isDebug" :event-id="eventId" />
        <div>
            <div v-if="!isLoading" class="controls">
                <button
                    :disabled="!isSeekableToNext"
                    @click="async () => {
                        currentEventIndex--;
                        await updateEventId(currentEventIndex);
                    }"
                >
                    <svg  xmlns="http://www.w3.org/2000/svg"  width="24"  height="24"  viewBox="0 0 24 24"  fill="none"  stroke="currentColor"  stroke-width="2"  stroke-linecap="round"  stroke-linejoin="round"  class="icon icon-tabler icons-tabler-outline icon-tabler-chevron-left"><path stroke="none" d="M0 0h24v24H0z" fill="none"/><path d="M15 6l-6 6l6 6" /></svg>
                </button>
                <button
                    :disabled="isLoading"
                    @click="async () => {
                        await shareCurrentURL();
                    }"
                >
                    <svg  xmlns="http://www.w3.org/2000/svg"  width="24"  height="24"  viewBox="0 0 24 24"  fill="none"  stroke="currentColor"  stroke-width="2"  stroke-linecap="round"  stroke-linejoin="round"  class="icon icon-tabler icons-tabler-outline icon-tabler-share"><path stroke="none" d="M0 0h24v24H0z" fill="none"/><path d="M6 12m-3 0a3 3 0 1 0 6 0a3 3 0 1 0 -6 0" /><path d="M18 6m-3 0a3 3 0 1 0 6 0a3 3 0 1 0 -6 0" /><path d="M18 18m-3 0a3 3 0 1 0 6 0a3 3 0 1 0 -6 0" /><path d="M8.7 10.7l6.6 -3.4" /><path d="M8.7 13.3l6.6 3.4" /></svg>
                </button>
                <button
                    :disabled="!isSeekableToPrev"
                    @click="async () => {
                        currentEventIndex++;
                        await updateEventId(currentEventIndex);
                    }"
                >
                    <svg  xmlns="http://www.w3.org/2000/svg"  width="24"  height="24"  viewBox="0 0 24 24"  fill="none"  stroke="currentColor"  stroke-width="2"  stroke-linecap="round"  stroke-linejoin="round"  class="icon icon-tabler icons-tabler-outline icon-tabler-chevron-right"><path stroke="none" d="M0 0h24v24H0z" fill="none"/><path d="M9 6l6 6l-6 6" /></svg>
                </button>
            </div>
        </div>
    </main>
</template>

<style scoped>
.controls {
    display: flex;
    justify-content: center;
    margin-top: 20px;

    button {
        margin: 0 10px;
        padding: 10px 20px;
        font-size: 16px;
        cursor: pointer;
    }
}
</style>
