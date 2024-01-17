module.exports = {
        preset: "ts-jest",
        testEnvironment: "jsdom", // 'jsdom' is suitable for React applications
        transform: {
                "^.+\\.(js|jsx|ts|tsx)$": "ts-jest", // Transform all JS/TS files with ts-jest
        },
        moduleNameMapper: {
                "\\.(css|less|sass|scss)$": "identity-obj-proxy", // Mock CSS imports (if needed)
                "\\.(gif|ttf|eot|svg)$": "identity-obj-proxy", // Mock image/font imports (if needed)
        },
        // Additional configuration...
};
