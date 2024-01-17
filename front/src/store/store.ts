// src/store.ts
import { configureStore } from "@reduxjs/toolkit";
import { useDispatch } from "react-redux";
import counterReducer from "./counterSlice";
import boardReducer from "./boardSlice";

export const store = configureStore({
        reducer: {
                counter: counterReducer,
                board: boardReducer,
        },
});

export type RootState = ReturnType<typeof store.getState>;
export type AppDispatch = typeof store.dispatch;
export const useAppDispatch = () => useDispatch<AppDispatch>();
