import {
  Box,
  Flex,
  Image,
  Link,
  Popover,
  PopoverArrow,
  PopoverBody,
  PopoverContent,
  PopoverTrigger,
  Text,
} from "@chakra-ui/react";
import cs from "clsx";
import { useLocation } from "react-router-dom";
import { NAV_ITEMS, NavItem } from "./menuConfig.ts";
import s from "./styles.module.scss";

const ProductMenu = ({ navItem }: { navItem: NavItem }) => {
  return (
    <Popover trigger="hover" key={navItem.label} placement="bottom-start">
      <PopoverTrigger>
        <Flex alignItems={"center"} gap="4px">
          <Text
            fontSize={"14px"}
            fontWeight={500}
            lineHeight={"110%"}
            letterSpacing={"0.03em"}
          >
            {navItem.label}
          </Text>
          <Image src={`/icons/ic-angle-down.svg`} />
        </Flex>
      </PopoverTrigger>
      <PopoverContent
        border="1px solid rgba(229, 231, 235, 0.40)"
        // borderBlock={'12px'}
        borderRadius={"12px"}
        boxShadow={"0px 0px 24px -6px rgba(0, 0, 0, 0.12)"}
        p="40px"
        w={"fit-content"}
      >
        <PopoverArrow />

        <PopoverBody
          p={0}
          display={"flex"}
          flexDirection={"column"}
          gap={"40px"}
        >
          <div className={s.dropdownItem}>
            <p className={s.dropdownItem_heading}>Agentic AI</p>
            <div className={s.dropdownItem_wrap}>
              <a
                href="https://eternalai.org/agents"
                rel="noopener noreferrer"
                className={s.dropdownItem_wrap_item}
              >
                <Image
                  src={"/icons_header/agent_launch.png"}
                  alt=""
                  width={20}
                  height={20}
                />
                <p>Agent Launchpad</p>
              </a>
              <a
                href="https://eternalai.org/agent-studio"
                rel="noopener noreferrer"
                className={s.dropdownItem_wrap_item}
              >
                <Image
                  src={"/icons_header/agent_studio.png"}
                  alt=""
                  width={20}
                  height={20}
                />
                <p>Agent Studio</p>
              </a>
              <a
                href="https://eternalai.org/labs"
                rel="noopener noreferrer"
                className={s.dropdownItem_wrap_item}
              >
                <Image
                  src={"/icons_header/agent_genomic.png"}
                  alt=""
                  width={20}
                  height={20}
                />
                <p>Agent Genomic Labs</p>
              </a>
            </div>
          </div>
          <div className={s.dropdownItem}>
            <p className={s.dropdownItem_heading}>NFT AI</p>
            <div className={s.dropdownItem_wrap}>
              <Link
                href="/nfts"
                rel="noopener noreferrer"
                className={s.dropdownItem_wrap_item}
              >
                <Image src={"/robot_01.png"} alt="" width={20} height={20} />
                <p>CryptoAgents</p>
              </Link>
              <Link
                href="/perceptrons"
                rel="noopener noreferrer"
                className={s.dropdownItem_wrap_item}
              >
                <Image
                  src={"/icons_header/perceptron.png"}
                  alt=""
                  width={20}
                  height={20}
                />
                <p>Perceptrons</p>
              </Link>
            </div>
          </div>
          <div className={s.dropdownItem}>
            <p className={s.dropdownItem_heading}>Physical Ai</p>
            <div className={s.dropdownItem_wrap}>
              <Link
                href="/neuron"
                rel="noopener noreferrer"
                className={s.dropdownItem_wrap_item}
              >
                <Image
                  src={"/icons_header/neuron.png"}
                  alt=""
                  width={20}
                  height={20}
                />
                <p>Neuron</p>
              </Link>
            </div>
          </div>
          <div className={s.dropdownItem}>
            <p className={s.dropdownItem_heading}>For Developers</p>
            <div className={s.dropdownItem_wrap}>
              <Link
                href="https://eternalai.org/api"
                rel="noopener noreferrer"
                className={s.dropdownItem_wrap_item}
              >
                <Image
                  src={"/icons_header/microscope.svg"}
                  alt=""
                  width={20}
                  height={20}
                />
                <p>API</p>
              </Link>
              <a
                href="https://github.com/eternalai-org/truly-open-ai"
                target="_blank"
                rel="noopener noreferrer"
                className={s.dropdownItem_wrap_item}
              >
                <Image
                  src={"/icons_header/note-text.svg"}
                  alt=""
                  width={20}
                  height={20}
                />
                <p>Open Source</p>
              </a>
            </div>
          </div>
        </PopoverBody>
      </PopoverContent>
    </Popover>
  );
};

type Props = {
  primaryColor?: "black" | "white";
  listNavItems?: Array<NavItem>;
};

const HeaderMenu = (props: Props) => {
  const location = useLocation();
  const pathName = location.pathname;
  console.log("stephen: pathname", pathName);

  return (
    <Flex
      className={s.deskMenu}
      style={{ "--color": props.primaryColor || "white" } as any}
    >
      {NAV_ITEMS?.map((navItem) => {
        return (
          <Link
            className={cs(
              navItem.rootHref
                ? pathName.includes(navItem.rootHref)
                  ? s.isActive
                  : ""
                : pathName === navItem.href
                ? s.isActive
                : "",
              s[props?.primaryColor || "white"]
            )}
            key={navItem.label}
            href={navItem.href ?? "#"}
            target={navItem.isNewWindow ? "_blank" : "_self"}
            color={props?.primaryColor || "white"}
          >
            <Box className={cs(s.borderLeft, s[navItem.className as any])} />
            <Image src={navItem.icon} />
          </Link>
        );
      })}
    </Flex>
  );
};

export default HeaderMenu;
