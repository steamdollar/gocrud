import { useSelector, useDispatch } from "react-redux";
import {
        increment,
        decrement,
        incrementByAmount,
        reset,
} from "../store/counterSlice";
import { RootState } from "../store/store";
import React, { useState } from "react";

export const Counter = () => {
        const count = useSelector((state: RootState) => state.counter.value);
        const dispatch = useDispatch();
        const [amount, setAmount] = useState(5);

        const handleInputChange = (e: React.ChangeEvent<HTMLInputElement>) => {
                setAmount(+e.target.value || 0);
        };

        return (
                <div>
                        <h1>Counter</h1>
                        <p>Count: {count}</p>
                        <button onClick={() => dispatch(increment())}>
                                Increment
                        </button>
                        <button onClick={() => dispatch(decrement())}>
                                Decrement
                        </button>

                        <button
                                onClick={() =>
                                        dispatch(incrementByAmount(amount))
                                }
                        >
                                Increment by{" "}
                                <input
                                        type="number"
                                        value={amount}
                                        onChange={handleInputChange}
                                />
                        </button>

                        <button onClick={() => dispatch(reset())}>reset</button>
                </div>
        );
};
