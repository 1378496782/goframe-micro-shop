import { defineStore } from 'pinia';
import {NoticeMessage, NoticeStates} from "/@/stores/interface";

export const noticeStore = defineStore({
    id: 'noticeStore',
    state: (): NoticeStates => <NoticeStates>({
        message: {}
    }),
    getters: {
        getMessages(): NoticeMessage {
            return this.message;
        },
    },
    actions: {
        setMessages(messages:NoticeMessage) {
            this.message = messages;
        },
    },
});
