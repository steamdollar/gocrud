import "./App.css";
import { BrowserRouter as Router, Route, Routes, Link } from "react-router-dom";
import { Tab, Tabs } from "@mui/material";
import { Counter } from "./components/Counter";

function App() {
        return (
                <div className="App">
                        <Router>
                                <Tabs>
                                        <Tab
                                                label="Counter"
                                                to="/"
                                                component={Link}
                                        />
                                </Tabs>
                                <Routes>
                                        <Route path="/" element={<Counter />} />
                                </Routes>
                        </Router>
                </div>
        );
}

export default App;
