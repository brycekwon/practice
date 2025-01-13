import type { Site, Metadata, Socials } from "@types";

export const SITE: Site = {
    TITLE: "Bryce Kwon - Computer Science Student",
    NAME: "Bryce Kwon",
    HEADER_PRIMARY: "Hi there! I'm a bug on the internet ☕️",
    DESCRIPTION_PRIMARY: `
        I'm a fourth-year undergraduate computer science student interested in software
        engineering, computer networks, and cybersecurity. This website is my personal
        corner on the internet where I can share my projects and experiences.
    `,
    HEADER_SECONDARY: "Let's connect",
    DESCRIPTION_SECONDARY: `
        If you want to get in touch with me about something or just want to say hi,
        feel free to send me an email or reach out on social media.
    `,
    NUM_POSTS_ON_HOMEPAGE: 3,
    NUM_PROJECTS_ON_HOMEPAGE: 3,
};

export const HOME: Metadata = {
    TITLE: "Home",
    DESCRIPTION: "Bryce Kwon, a computer science student and aspiring software engineer. Explore my some of my projects, interests, and professional history.",
};

export const POSTS: Metadata = {
    TITLE: "Posts",
    DESCRIPTION: "A collection of posts on topics I am passionate about.",
};

export const PROJECTS: Metadata = {
    TITLE: "Projects",
    DESCRIPTION: "A collection of my projects, with links to repositories and demos.",
};

export const SOCIALS: Socials = [
    {
        NAME: "email",
        HREF: "mailto:me@brycekwon.com",
    },
    { 
        NAME: "github",
        HREF: "https://github.com/brycekwon"
    },
    { 
        NAME: "linkedin",
        HREF: "https://www.linkedin.com/in/brycekwon",
    },
];
