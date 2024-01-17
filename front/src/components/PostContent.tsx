import React, { useEffect, useState } from "react";
import { useParams } from "react-router-dom";

export const PostContentComponent = () => {
        const { idx } = useParams();
        const [postIndex, setPostIndex] = useState("");

        useEffect(() => {}, [idx]);

        return (
                <div>
                        <h1>Content</h1>
                        <p>{postIndex}</p>
                </div>
        );
};
