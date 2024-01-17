import React, { useEffect } from "react";
import { useSelector } from "react-redux";
import { useAppDispatch } from "../store/store";
import { Post, fetchBoardData } from "../store/boardSlice";
import { RootState } from "../store/store";
import { Link } from "react-router-dom";

export const Board = () => {
        const dispatch = useAppDispatch();

        const boardData = useSelector((state: RootState) => state.board);

        useEffect(() => {
                if (boardData.status === "idle") {
                        dispatch(fetchBoardData());
                        console.log(boardData);
                }
        }, [boardData.status, dispatch]);

        const postsInfo = (boardInfo: Post[]): any => {
                return boardInfo.map((v, i) => {
                        return (
                                <div key={i}>
                                        <span style={{ marginRight: "50px" }}>
                                                {v.Idx}
                                        </span>
                                        <span style={{ marginRight: "50px" }}>
                                                {v.Author}
                                        </span>

                                        <span style={{ marginRight: "50px" }}>
                                                {v.Date}
                                        </span>

                                        <span style={{ marginRight: "50px" }}>
                                                {v.Views}
                                        </span>
                                        <Link to={`/post/${v.Idx}`}>
                                                Read more
                                        </Link>
                                </div>
                        );
                });
        };

        return (
                <div>
                        {" "}
                        <span>{postsInfo(boardData.data)}</span>{" "}
                </div>
        );
};
