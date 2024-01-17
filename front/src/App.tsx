import "./App.css";
import React, { ChangeEvent, useState } from "react";
import { BrowserRouter as Router, Route, Routes, Link } from "react-router-dom";
import { Tab, Tabs } from "@mui/material";
import { Counter } from "./components/Counter";
import { Board } from "./components/Board";
import { PostContentComponent } from "./components/PostContent";

function App() {
        const [value, setValue] = useState<number>(0);

        const handleChange = (
                event: React.SyntheticEvent<{}>,
                newValue: number
        ) => {
                setValue(newValue);
                // Additional logic for tab change...
        };

        return (
                <div className="App">
                        <Router>
                                <Tabs value={value} onChange={handleChange}>
                                        <Tab
                                                label="Counter"
                                                to="/"
                                                component={Link}
                                                value={0}
                                        />
                                        <Tab
                                                label="Board"
                                                to="/board"
                                                component={Link}
                                                value={1}
                                        />
                                </Tabs>
                                <Routes>
                                        <Route path="/" element={<Counter />} />
                                        <Route
                                                path="/board"
                                                element={<Board />}
                                        />
                                        <Route
                                                path="/post/:idx"
                                                element={
                                                        <PostContentComponent />
                                                }
                                        />
                                </Routes>
                        </Router>
                </div>
        );
}

export default App;
