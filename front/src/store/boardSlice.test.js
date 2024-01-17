import { fetchBoardData } from "./boardSlice";
import { configureStore } from "@reduxjs/toolkit";
import boardReducer from "./boardSlice";

describe("board slice integration test", () => {
        let store;

        beforeAll(() => {
                store = configureStore({ reducer: { board: boardReducer } });
        });

        it("fetch data from backend", async () => {
                await store.dispatch(fetchBoardData());

                const state = store.getstate().board;
                console.log(state);
        });
});
