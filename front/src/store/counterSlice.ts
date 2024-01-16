import { createSlice, PayloadAction } from "@reduxjs/toolkit";

interface Counter {
        value: number;
}

const initialState: Counter = {
        value: 0,
};

// `createSlice`를 이용해 slice를 생성
export const counterSlice = createSlice({
        name: "counter",
        initialState,
        reducers: {
                increment: (state) => {
                        state.value += 1;
                },

                decrement: (state) => {
                        state.value -= 1;
                },
                reset: (state) => {
                        state.value = 0;
                },
                incrementByAmount: (state, action: PayloadAction<number>) => {
                        state.value += action.payload;
                },
        },
});

// 생성된 액션, 리듀서를 export
// reducer가 store에 import 되어야 한다.
export const { increment, decrement, incrementByAmount, reset } =
        counterSlice.actions;

// 이 reducer는 앞서 정의한 함수들이 아님.
// createSlice 함수에 의해 자동 생성된 함수임.
// 이 함수는 slice 내부에서 정의된 모든 액션들을 다룰 수 있고, 상태를 업데이트 할 수 있다.
//
// reducers 내에 생성된 increment, decrement 등의 함수는 `action creators` 라고 하며,
// store에 디스패치 될 액션들을 생성한다.
// export default 이므로 import 할때는 원하는 이름을 붙여줄 수 있다.
export default counterSlice.reducer;
