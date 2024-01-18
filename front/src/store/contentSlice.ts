import { createAsyncThunk, createSlice, PayloadAction } from "@reduxjs/toolkit";
import axios from "axios";

export interface PostContent {
        Idx: number | null;
        Content: string;
        status: string;
}

const initialState: PostContent = {
        Idx: null,
        Content: "",
        status: "idle",
};

export const fetchPostContent = createAsyncThunk(
        "posts/fetchById",
        async (idx, thunkAPI) => {
                const response = await axios.get(
                        `http://localhost:4000/getPostContent?idx=${idx}`
                );
                console.log(response.data);
                return response.data;
        }
);

export const postSlice = createSlice({
        name: "posts",
        initialState,
        reducers: {},
        extraReducers: (stat) => {
                stat.addCase(fetchPostContent.pending, (state) => {
                        state.status = "loading";
                })
                        .addCase(
                                fetchPostContent.fulfilled,
                                (state, action) => {
                                        state.status = "succeeded";
                                        state.Content = action.payload;
                                }
                        )
                        .addCase(fetchPostContent.rejected, (state, action) => {
                                state.status = "failed";
                        });
        },
});

export default postSlice.reducer;
