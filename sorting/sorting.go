package sorting

//Sorting is an interface for sorting list
type Sorting interface {
	Sorting(list []int)
}

type quickSortingImpl struct{}

func (s *quickSortingImpl) Sorting(a []int) {
	qSort(a, 0, len(a)-1)
}
func qSort(in []int, low int, high int) {
	if low >= high {
		return
	}
	key := in[low]
	i, j := low+1, high
	for {
		for key > in[i] {
			if i == high {
				break
			}
			i++
		}
		for key < in[j] {
			if j == low {
				break
			}
			j--
		}
		if i >= j {
			break
		}
		in[i], in[j] = in[j], in[i]
		i++
		j--
	}
	in[low], in[j] = in[j], in[low]
	qSort(in, low, j-1)
	qSort(in, j+1, high)
}

type bubblingSortingImpl struct{}

func (s *bubblingSortingImpl) Sorting(a []int) {

}

var (
	qSorting = new(quickSortingImpl)
)

//QSorting is a quick sorting algorithm
func QSorting() Sorting {
	return qSorting
}
