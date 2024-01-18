import React, { useEffect, useState } from "react";
import { useDispatch } from "react-redux";
import { useParams } from "react-router-dom";

export const PostContentComponent = () => {
        const { idx } = useParams();
        const dispatch = useDispatch();
        // const postContent = useSelector((state) => state.board.postContent);

        useEffect(() => {}, [idx, dispatch]);

        return (
                <div>
                        <h1>Content</h1>
                        <p>{idx}</p>
                </div>
        );
};
