import { createAsyncThunk, createSlice, PayloadAction } from "@reduxjs/toolkit";
import axios from "axios";

export interface Post {
        Idx: number;
        Author: string;
        Date: any;
        Content: string;
        Views: number;
}

interface BoardState {
        data: Post[];
        status: string;
}

const initialState: BoardState = {
        data: [],
        status: "idle",
};

export const fetchBoardData = createAsyncThunk("board/fetchData", async () => {
        const response = await axios.get("http://localhost:4000/sample");
        return response.data;
});

export const boardSlice = createSlice({
        name: "board",
        initialState,
        reducers: {},
        extraReducers: (stat) => {
                stat.addCase(fetchBoardData.pending, (state) => {
                        state.status = "loading";
                })
                        .addCase(fetchBoardData.fulfilled, (state, action) => {
                                state.status = "succeeded";
                                state.data = action.payload;
                        })
                        .addCase(fetchBoardData.rejected, (state, action) => {
                                state.status = "failed";
                        });
        },
});

export default boardSlice.reducer;
