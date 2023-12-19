# CLI Applcation With Cobra
cobra란 명령, 인수 및 플래그 구조 기반의 CLI를 제작을 도와주는 툴이다.

# cobra.Command
cobra.Command 구조체는 명령어 하나를 생성할 때 필연적으로 생성해줘 하는 구조체이다. 즉 명령어를 생성하기 위해선 cobra.Command 구조체를 사용해야 한다.

cobra.Command 구조체의 구성요소들은 아주 많다. 내가 오늘 사용한 요소들에 대해 설명해주겠다.
- Use: 말그대로 사용하는 명령어
- Short: 간단한 설명
- Long: 긴 설명
- Run: 이 명령어의 비지니스 로직
- Aliases: 별칭
- Args: 에상 인자 위치 (값으로 cobra.ExactArgs로 사용한다.)

rootCmd.AddCommand 함수를 이용해 command를 추가할 수 있다.

# cobra flag
flag란? 특정 커맨드의 추가 옵션을 넣음으로써 기능을 다체롭게 사용하게 해서 용이하게 CLI를 쓰게 해주는 기능 

cobra에서는 플래그가 총 두 가지가 있다.
- 로컬 flag : 지정한 커맨드에서만 동작하는 커맨드, 마치 지역변수와 같다.
- 지속 flag : 모든 커맨드에 적용되는 커맨드, 마치 전역변수와 같다.

rootCommand의 Version 요소를 지정해 줄 경우
--version 치면 해당 CLI의 버전이 출력된다.

로컬 flag 같은 경우 해당 플래그의 command.Flags.BoolVarP(p, name, shorthead, value, usage)함수를 통해서 사용할 수 있다.