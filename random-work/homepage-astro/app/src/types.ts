export type Site = {
    TITLE: string,
    NAME: string;
    HEADER_PRIMARY: string;
    DESCRIPTION_PRIMARY: string;
    HEADER_SECONDARY: string;
    DESCRIPTION_SECONDARY: string;
    NUM_POSTS_ON_HOMEPAGE: number;
    NUM_PROJECTS_ON_HOMEPAGE: number;
};

export type Metadata = {
    TITLE: string;
    DESCRIPTION: string;
};

export type Socials = {
    NAME: string;
    HREF: string;
}[];
