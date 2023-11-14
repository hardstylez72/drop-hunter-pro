import {Network} from "@/generated/flow";
import {ProfileType} from "@/generated/profile";
import {Airdrop, TaskJob, TaskSpec, Universal} from "@/components/tasks/utils";
import TaskDeployStarkNetAccount from "@/components/tasks/OTHER/StarknetDeployAccount/Block.vue";
import DeployStarkNetAccount from "@/components/tasks/OTHER/StarknetDeployAccount/Menu.vue";
import StarkNetId from "@/components/tasks/OTHER/StarknetId/Block.vue";
import MenuStarkNetId from "@/components/tasks/OTHER/StarknetId/Menu.vue";

export const StarknetIdMintSpec: TaskSpec = {
  deprecated: false,
  canBeEstimated: true,
  descFn(task) {
    const p = task.starkNetIdMintTask
    return ''
  },
  component: StarkNetId,
  menu: MenuStarkNetId,
  service: {
    name: 'StarkNetId',
    link: 'https://app.starknet.id/',
    img: '/icons/starknetid.svg',
    op: 'mint uniq id',
  },
  job: TaskJob.NFT,
  networks: new Set<Network>([Network.StarkNet]),
  airdrops: new Set<Airdrop>([Airdrop.StarkNet]),
  profileType: new Set([ProfileType.StarkNet])
}
